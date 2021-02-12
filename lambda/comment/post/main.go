package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/knoebber/aws-utils/lambda"
	"nicolasknoebber.com/comment"
)

// HandleRequest saves a comment to DynamoDB.
func HandleRequest(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	var c comment.Comment

	response.Headers = map[string]string{"Access-Control-Allow-Origin": "*"}
	svc := dynamodb.New(session.Must(session.NewSession()))

	if err = json.Unmarshal([]byte(request.Body), &c); err != nil {
		response.StatusCode = 400
		return
	}

	c.Timestamp = time.Now().Unix() * 1000 // For JavaScript Date object.

	if err = c.Save(svc); err != nil {
		response.StatusCode = 500
		return
	}

	lambda.SetResponseBody(&response, c)
	return
}

func main() {
	lambda.Start(HandleRequest)
}
