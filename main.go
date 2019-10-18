package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/markrtucker/chorehound/chores"
)

// HandleRequest TODO godoc
func HandleRequest(ctx context.Context) (string, error) {
	return fmt.Sprintf("Hello from AWS!"), nil
}

func main() {
	lambda.Start(HandleRequest)
}

// TODO: See https://ewanvalentine.io/how-im-writing-serverless-services-in-golang-these-days/

// Use cases:
// Create a new chore and store it in bucket (S3/OS??)
// Given a chore ID, read the chore JSON from the bucket (TBD: Why?)
// Given a chore, mark it as done (advance the schedule)
// Given a chore, hound the person scheduled to do it

// TODO: Is each use case a separate lambda function??

// NewChore TODO godoc
func NewChore(ctx context.Context, c chores.Chore) error {

	// TODO: basedir location?
	repo := chores.NewFileSystemRepository("/var/chores/")
	c.SaveTo(ctx, repo)

	return nil
}
