package microblog

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const bucketname = "nicolasknoebber.com"

// Post is a microblog post.
type Post struct {
	Title  string  `json:"title" dynamodbav:"title"`
	Body   string  `json:"body" dynamodbav:"body"`
	Images []image `json:"images" dynamodbav:"images"`
}

// Image is an image for a post.
// They are stored in an S3 Bucket.
type Image struct {
	Title    string `json:"title" dynamodbav:"title"`
	Filename string `json:"filename" dynamodbav:"filename"`
	Alt      string `json:"alt" dynamodbav:"alt"`
}
