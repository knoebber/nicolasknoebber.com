package main

import (
 "bytes"
 "fmt"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-lambda-go/lambda"
 "github.com/aws/aws-sdk-go/service/s3"
 "github.com/aws/aws-sdk-go/aws/session"
)

// Request is the structure that HandleRequest expects to receive.
type Request struct {
	Message string `json:"message"`
}

// HandleRequest processes a lambda request.
func HandleRequest(r Request) (string, error) {
  buffer, err := draw()
  if err != nil {
    return "", err
  }

  // Create a S3 client
  session := session.Must(session.NewSession())
  svc := s3.New(session)

  reader := bytes.NewReader(buffer.Bytes())
  fmt.Printf("%d bytes",reader.Len())

  putInput := s3.PutObjectInput{
    Bucket: aws.String("nicolasknoebber.com"),
    Body:   reader,
    Key:    aws.String("test_upload.png"),
  }

  result, err := svc.PutObject(&putInput)
  if err != nil {
    fmt.Println(err.Error)
    return "", err
  }

	return fmt.Sprintf("put object result: %s", result), nil
}


var dev = true
func main() {
  if !dev {
    lambda.Start(HandleRequest)
  }
  draw()
}
