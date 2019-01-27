package main

import (
	"fmt"
	"strconv"

	"github.com/ldmtam/serverless-todo/src/database"
	"github.com/ldmtam/serverless-todo/src/repositories"
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

	mysql, err := database.NewMySQL()
	if err != nil {
		return utils.ResponseError(err)
	}

	todoRepo := repositories.NewTodoRepo(mysql)
	err = todoRepo.DeleteTodoWithID(id)
	if err != nil {
		return utils.ResponseError(err)
	}

	return utils.ResponseOK(fmt.Sprintf("Successfully delete todo with ID %v", id))
}

func main() {
	lambda.Start(Handler)
}
