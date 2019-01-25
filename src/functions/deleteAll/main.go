package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverless-todo/src/models"
	"github.com/serverless-todo/src/utils"
)

// Handler ...
func Handler() (events.APIGatewayProxyResponse, error) {
	utils.MockTodos = []models.Todo{}

	return utils.ResponseOK(utils.MockTodos)
}

func main() {
	lambda.Start(Handler)
}