package microblog

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Post is a microblog post.
type Post struct {
	Hash    string    `json:"id" dynamodbav:"id"`
	Created time.Time `json:"created" dynamodbav:"created"`
	Text    string    `json:"body" dynamodbav:"body"`
	Images  []Image   `json:"images" dynamodbav:"images"`
}

func (p *Post) primaryKey() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{"hash": {S: aws.String(p.Hash)}}
}

// Get gets a post by its primary key, hash.
func (p *Post) Get(svc *dynamodb.DynamoDB) error {
	output, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       p.primaryKey(),
	})
	if err != nil {
		return fmt.Errorf("getting dynamodb entry for %q: %w", p.Hash, err)
	}

	if output.Item == nil {
		return ErrPostNotFound
	}

	err = dynamodbattribute.UnmarshalMap(output.Item, p)
	if err != nil {
		return fmt.Errorf("unmarshaling post %q: %w", p.Hash, err)
	}

	return nil
}

// Save saves the post to dynamodb.
func (p *Post) Save(svc *dynamodb.DynamoDB) error {
	item, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return fmt.Errorf("marshalling post: %w", err)
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("saving post: %w", err)
	}

	return nil
}
