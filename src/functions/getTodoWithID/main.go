package main

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/serverless-todo/src/models"
	"github.com/serverless-todo/src/utils"
)

// Handler ...
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ResponseError(err)
	}

	var returnedData models.Todo
	for _, todo := range utils.MockTodos {
		if todo.ID == id {
			returnedData = todo
			break
		}
	}

	return utils.ResponseOK(returnedData)
}

func main() {
	lambda.Start(Handler)
}
