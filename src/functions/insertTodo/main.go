package main

import (
	"encoding/json"

	"github.com/ldmtam/serverless-todo/src/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ldmtam/serverless-todo/src/models"
)

// Handler ...
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var todo models.Todo

	err := json.Unmarshal([]byte(request.Body), &todo)
	if err != nil {
		return utils.ResponseError(err)
	}

	utils.MockTodos = append(utils.MockTodos, todo)

	return utils.ResponseOK(utils.MockTodos)
}

func main() {
	lambda.Start(Handler)
}
