package main

import (
	"strconv"

	"github.com/ldmtam/serverless-todo/src/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler ...
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ResponseError(err)
	}

	for index, todo := range utils.MockTodos {
		if todo.ID == id {
			left := utils.MockTodos[:index]
			right := utils.MockTodos[index+1:]
			utils.MockTodos = append(left, right...)
		}
	}

	return utils.ResponseOK(utils.MockTodos)
}

func main() {
	lambda.Start(Handler)
}
