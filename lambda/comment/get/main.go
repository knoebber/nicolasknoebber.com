package main

import (
	"errors"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/knoebber/aws-utils/lambda"
	"nicolasknoebber.com/comment"
)

// HandleRequest returns a list of comments for a post.
func HandleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	var result []comment.Comment

	response.Headers = map[string]string{"Access-Control-Allow-Origin": "*"}
	postNumber, _ := strconv.Atoi(request.PathParameters["post"])
	if postNumber < 1 {
		err = errors.New("positive poster number is required")
		response.StatusCode = 400
	}

	svc := dynamodb.New(session.Must(session.NewSession()))
	result, err = comment.List(svc, postNumber)
	if err != nil {
		response.StatusCode = 500
		return
	}

	lambda.SetResponseBody(&response, result)
	response.StatusCode = 200
	return
}

func main() {
	lambda.Start(HandleRequest)
}
