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
