package microblog

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	bucketName = "nicolasknoebber.com"
	tableName  = "microblog"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrImageNotFound = errors.New("image not found")
)

// Used for generating random IDs.
func randomBytes(n int) ([]byte, error) {
	buff := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, buff); err != nil {
		return nil, err
	}

	return buff, nil
}

func ListPosts(sess *session.Session) ([]Post, error) {
	var result []Post

	db := dynamodb.New(sess)
	output, err := db.Query(&dynamodb.QueryInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, fmt.Errorf("listing posts: %w", err)
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(output.Items, &result); err != nil {
		return nil, fmt.Errorf("unmarshalling comments: %w", err)
	}

	return result, nil
}

func GetPost(sess *session.Session, postID string) (*Post, error) {
	result := &Post{ID: postID}

	if err := result.Get(dynamodb.New(sess)); err != nil {
		return nil, err
	}

	return result, nil
}

func CreatePost(sess *session.Session, text string) (*Post, error) {
	db := dynamodb.New(sess)

	randomBuff, err := randomBytes(7)
	if err != nil {
		return nil, fmt.Errorf("generating random bytes for post id: %w", err)
	}

	post := &Post{
		ID:      base64.URLEncoding.EncodeToString(randomBuff),
		Created: time.Now().UTC(),
		Text:    text,
		Images:  []Image{},
	}

	if err := post.Save(db); err != nil {
		return nil, err
	}

	return post, nil
}

func AttachImage(sess *session.Session, postID string, img *Image) (*Post, error) {
	db := dynamodb.New(sess)
	s3 := s3.New(sess)

	post := &Post{ID: postID}
	if err := post.Get(db); err != nil {
		return nil, err
	}

	if err := img.Upload(s3, postID); err != nil {
		return nil, err
	}

	post.Images = append(post.Images, *img)

	if err := post.Save(db); err != nil {
		return post, nil
	}

	return post, nil
}

func UpdatePost(sess *session.Session, postID string, newText string) (*Post, error) {
	db := dynamodb.New(sess)

	post := &Post{ID: postID}
	if err := post.Get(db); err != nil {
		return nil, err
	}

	post.Text = newText

	if err := post.Save(db); err != nil {
		return post, nil
	}

	return post, nil

}

func UpdateImage(sess *session.Session, postID, filename, newCaption, newAlt string) (*Post, error) {
	db := dynamodb.New(sess)

	post := &Post{ID: postID}
	if err := post.Get(db); err != nil {
		return nil, err
	}

	for i := range post.Images {
		img := post.Images[i]

		if img.Filename == filename {
			post.Images[i] = Image{
				Caption: newCaption,
				Alt:     newAlt,
			}
			break
		}

		if i == len(post.Images)-1 {
			return nil, ErrImageNotFound
		}
	}

	if err := post.Save(db); err != nil {
		return post, nil
	}

	return post, nil
}

func DeleteImage(sess *session.Session, postID, filename string) (*Post, error) {
	db := dynamodb.New(sess)
	s3 := s3.New(sess)

	post := &Post{ID: postID}
	if err := post.Get(db); err != nil {
		return nil, err
	}

	newImages := []Image{}
	imageToDelete := Image{}
	for _, img := range post.Images {
		if img.Filename == filename {
			imageToDelete = img
			continue
		}

		newImages = append(newImages, img)
	}
	if imageToDelete.Filename == "" {
		return nil, ErrImageNotFound
	}

	if err := imageToDelete.Delete(s3, postID); err != nil {
		return nil, err
	}

	return post, nil
}

func DeletePost(sess *session.Session, postID string) error {
	db := dynamodb.New(sess)
	s3 := s3.New(sess)

	post := &Post{ID: postID}
	if err := post.Get(db); err != nil {
		return err
	}

	for _, img := range post.Images {
		if err := img.Delete(s3, postID); err != nil {
			return err
		}
	}

	if _, err := db.DeleteItem(&dynamodb.DeleteItemInput{Key: post.primaryKey()}); err != nil {
		return fmt.Errorf("deleting post %q from dynamodb: %w", postID, err)
	}

	return nil
}
