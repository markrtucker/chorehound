package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
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
// Given a chore ID, read the chore JSON from the bucket
// Given a chore, mark it as done (advance the schedule)
// Given a chore, hound the person scheduled to do it
