package utils

import "github.com/ldmtam/serverless-todo/src/models"

// MockTodos todo mock data
var MockTodos = []models.Todo{
	{
		ID:         1,
		Content:    "Learn AWS Lambda",
		IsFinished: false,
	},
	{
		ID:         2,
		Content:    "Lose weight",
		IsFinished: false,
	},
}
