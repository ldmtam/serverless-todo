package repositories

import (
	"github.com/ldmtam/serverless-todo/src/models"
)

// TodoRepo ...
type TodoRepo struct {
	database Database
}

// NewTodoRepo ...
func NewTodoRepo(db Database) *TodoRepo {
	return &TodoRepo{database: db}
}

// InsertTodo ...
func (todoRepo *TodoRepo) InsertTodo(todo models.Todo) error {
	return todoRepo.database.InsertTodo(todo)
}

// GetAllTodos ...
func (todoRepo *TodoRepo) GetAllTodos() ([]models.Todo, error) {
	return todoRepo.database.GetAllTodos()
}

// GetTodoWithID ...
func (todoRepo *TodoRepo) GetTodoWithID(id int) (models.Todo, error) {
	return todoRepo.database.GetTodoWithID(id)
}

// DeleteTodoWithID ...
func (todoRepo *TodoRepo) DeleteTodoWithID(id int) error {
	return todoRepo.database.DeleteTodoWithID(id)
}

// DeleteAllTodos ...
func (todoRepo *TodoRepo) DeleteAllTodos() error {
	return todoRepo.database.DeleteAllTodos()
}
