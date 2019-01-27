package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ldmtam/serverless-todo/src/database"
	"github.com/ldmtam/serverless-todo/src/repositories"
	"github.com/ldmtam/serverless-todo/src/utils"
)

// Handler ...
func Handler() (events.APIGatewayProxyResponse, error) {
	mysql, err := database.NewMySQL()
	if err != nil {
		return utils.ResponseError(err)
	}

	todoRepo := repositories.NewTodoRepo(mysql)
	err = todoRepo.DeleteAllTodos()
	if err != nil {
		return utils.ResponseError(err)
	}

	return utils.ResponseOK("Successfully delete all")
}

func main() {
	lambda.Start(Handler)
}
