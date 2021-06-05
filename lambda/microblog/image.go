package microblog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Image is an image for a post.
// They are stored in an S3 Bucket.
type Image struct {
	Filename string `json:"filename" dynamodbav:"filename"`
	Caption  string `json:"caption" dynmodbav:"caption"`
	Alt      string `json:"alt" dynamodbav:"alt"`

	Data io.Reader `json:"-" dynamodbav:"-"`
}

func (i *Image) Key(postID string) string {
	return fmt.Sprintf("/microblog/images/%s/%s", postID, i.Filename)
}

// Upload uploads an image to s3.
func (i *Image) Upload(svc *s3.S3, postID string) error {
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(i.Key(postID)),
		Body:   i.Data,
	}

	uploader := s3manager.NewUploaderWithClient(svc)

	// Perform an upload.
	_, err := uploader.Upload(upParams)
	if err != nil {
		return fmt.Errorf("saving image %q for post %q: %w", i.Filename, postID, err)
	}

	return nil
}

func (i *Image) Delete(svc *s3.S3, postID string) error {
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(i.Key(postID)),
	})
	if err != nil {
		return fmt.Errorf("deleting image %q for post %q: %w", i.Filename, postID, err)
	}

	return nil
}

// UnmarshalImage unmarshals an image from json in body.
// The first part is a json encoded Image, the second part is binary data for the image.
func UnmarshalImage(body io.Reader, boundry string) (*Image, error) {
	mr := multipart.NewReader(body, boundry)

	jsonPart, err := mr.NextPart()
	if err != nil {
		return nil, fmt.Errorf("getting json part: %w", err)
	}

	if jsonPart.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("expected json part to be content type application/json")
	}

	result := new(Image)
	if err := json.NewDecoder(jsonPart).Decode(result); err != nil {
		return nil, fmt.Errorf("decoding image json: %w", err)
	}

	if err := jsonPart.Close(); err != nil {
		return nil, fmt.Errorf("closing json part: %w", err)
	}

	result.Data, err = mr.NextPart()
	if err != nil {
		return nil, fmt.Errorf("getting image data part: %w", err)
	}

	return result, nil
}
