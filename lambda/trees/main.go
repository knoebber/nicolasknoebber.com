package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// TreeParam is the structure that HandleRequest expects to receive.
type TreeParam struct {
	LeftLength  float64 `json:"leftLength"`
	LeftAngle   float64 `json:"leftAngle"`
	RightLength float64 `json:"rightLength"`
	RightAngle  float64 `json:"rightAngle"`
}

// HandleRequest processes a lambda request.
func HandleRequest(p TreeParam) (string, error) {
	buffer, err := createTree(p)
	if err != nil {
		return "", err
	}

	// Create a S3 client
	session := session.Must(session.NewSession())
	svc := s3.New(session)

	reader := bytes.NewReader(buffer.Bytes())
	fmt.Printf("%d bytes", reader.Len())

	putInput := s3.PutObjectInput{
		Bucket: aws.String("nicolasknoebber.com"),
		Body:   reader,
		Key:    aws.String("lambda-go-tree.png"),
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
	// For testing locally
	if dev {
		createTree(TreeParam{LeftLength: 14, RightLength: 12, LeftAngle: 12, RightAngle: 45})
		return
	}
	lambda.Start(HandleRequest)
}
