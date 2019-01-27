package repositories

import (
	"github.com/ldmtam/serverless-todo/src/models"
)

// Database ...
type Database interface {
	InsertTodo(todo models.Todo) error
	GetAllTodos() ([]models.Todo, error)
	GetTodoWithID(id int) (models.Todo, error)
	DeleteTodoWithID(id int) error
	DeleteAllTodos() error
}
