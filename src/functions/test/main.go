package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverless-todo/src/utils"
)

// Handler ...
func Handler() (events.APIGatewayProxyResponse, error) {
	return utils.ResponseOK("This is a test function. It should be returned :)")
}

func main() {
	lambda.Start(Handler)
}
