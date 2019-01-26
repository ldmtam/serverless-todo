package main

import (
	"encoding/json"

	"github.com/ldmtam/serverless-todo/src/repositories"

	"github.com/ldmtam/serverless-todo/src/database"

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

	mysql, err := database.NewMySQL()
	if err != nil {
		return utils.ResponseError(err)
	}

	todoRepo := repositories.NewTodoRepo(mysql)
	todoRepo.InsertTodo(todo)

	todoList, err := todoRepo.GetAllTodos()
	if err != nil {
		return utils.ResponseError(err)
	}

	return utils.ResponseOK(todoList)
}

func main() {
	lambda.Start(Handler)
}
