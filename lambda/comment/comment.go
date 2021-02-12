package comment

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "comment"

// Comment is a blog comment.
// PostNumber and Timestamp make up a composite key.
type Comment struct {
	PostNumber int    `dynamodbav:"post_number" json:"postNumber"`
	Timestamp  int64  `dynamodbav:"time_stamp" json:"timestamp"`
	Name       string `dynamodbav:"comment_name" json:"commentName"`
	Body       string `dynamodbav:"comment_body" json:"commentBody"`
}

// Save saves the comment.
func (c Comment) Save(svc *dynamodb.DynamoDB) error {
	if c.PostNumber < 1 {
		return fmt.Errorf("post number is required")
	}
	if c.Timestamp < 1 {
		return fmt.Errorf("timestamp is required")
	}
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}

	if c.Body == "" {
		return fmt.Errorf("body is required")
	}

	item, err := dynamodbattribute.MarshalMap(c)
	if err != nil {
		return fmt.Errorf("marshaling %+v: %w", c, err)
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("putting comment %+v: %w", c, err)
	}

	return nil
}

// List retrieves all comments for the post number.
func List(svc *dynamodb.DynamoDB, postNumber int) ([]Comment, error) {
	var result []Comment

	output, err := svc.Query(&dynamodb.QueryInput{
		ScanIndexForward: aws.Bool(true),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v1": {N: aws.String(strconv.Itoa(postNumber))},
		},
		KeyConditionExpression: aws.String("post_number = :v1"),
		TableName:              aws.String(tableName),
	})

	if err != nil {
		return nil, fmt.Errorf("listing comments for post %d", postNumber)
	}

	if err := dynamodbattribute.UnmarshalListOfMaps(output.Items, &result); err != nil {
		return nil, fmt.Errorf("unmarshalling comments: %w", err)
	}

	return result, nil
}
