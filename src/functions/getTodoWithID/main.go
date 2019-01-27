package main

import (
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ldmtam/serverless-todo/src/database"
	"github.com/ldmtam/serverless-todo/src/repositories"
	"github.com/ldmtam/serverless-todo/src/utils"
)

// Handler ...
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return utils.ResponseError(err)
	}

	mysql, err := database.NewMySQL()
	if err != nil {
		return utils.ResponseError(err)
	}

	todoRepo := repositories.NewTodoRepo(mysql)

	todo, err := todoRepo.GetTodoWithID(id)
	if err != nil {
		return utils.ResponseError(err)
	}

	return utils.ResponseOK(todo)
}

func main() {
	lambda.Start(Handler)
}
