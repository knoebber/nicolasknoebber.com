package microblog

import (
	"fmt"
	"io"

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
}

func (i *Image) Key(postHash string) string {
	return fmt.Sprintf("/microblog/images/%s/%s", postHash, i.Filename)
}

// Upload uploads an image to s3.
func (i *Image) Upload(svc *s3.S3, postHash string, body io.Reader) error {
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(i.Key(postHash)),
		Body:   body,
	}
	uploader := s3manager.NewUploaderWithClient(svc)

	// Perform an upload.
	_, err := uploader.Upload(upParams)
	if err != nil {
		return fmt.Errorf("saving image %q for post %q: %w", i.Filename, postHash, err)
	}

	return nil
}

func (i *Image) Delete(svc *s3.S3, postHash string) error {
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(i.Key(postHash)),
	})
	if err != nil {
		return fmt.Errorf("deleting image %q for post %q: %w", i.Filename, postHash, err)
	}

	return nil
}
