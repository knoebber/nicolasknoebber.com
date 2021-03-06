package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var dev = false

// TreeParam is the structure that HandleRequest expects to receive.
type TreeParam struct {
	LeftLength  float64 `json:"leftLength"`
	LeftAngle   float64 `json:"leftAngle"`
	RightLength float64 `json:"rightLength"`
	RightAngle  float64 `json:"rightAngle"`
}

// HandleRequest processes a lambda request.
func HandleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	response.Headers = map[string]string{"Access-Control-Allow-Origin": "*"}

	var (
		p      TreeParam
		buffer *bytes.Buffer
	)

	if err = json.Unmarshal([]byte(request.Body), &p); err != nil {
		response.StatusCode = 400
		return
	}
	buffer, err = createTree(p)
	if err != nil {
		response.StatusCode = 500
		return
	}

	fileName := fmt.Sprintf("lambda-go-tree-%d.png", time.Now().Unix())
	// Create a S3 client
	session := session.Must(session.NewSession())
	svc := s3.New(session)

	reader := bytes.NewReader(buffer.Bytes())
	putInput := s3.PutObjectInput{
		Bucket: aws.String("nicolasknoebber.com"),
		Body:   reader,
		Key:    aws.String(fmt.Sprintf("/posts/images/trees/%s", fileName)),
	}

	_, err = svc.PutObject(&putInput)
	if err != nil {
		response.StatusCode = 500
		return
	}

	response.StatusCode = 200
	response.Body = fmt.Sprintf(`{"message":"%s"}`, fileName)
	return
}

func main() {
	// For testing locally
	if dev {
		createTree(TreeParam{LeftLength: 14, RightLength: 12, LeftAngle: 12, RightAngle: 45})
		return
	}
	lambda.Start(HandleRequest)
}
