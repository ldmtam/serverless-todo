package database

import (
	"time"

	"github.com/ldmtam/serverless-todo/src/models"
)

// InsertTodo ...
func (mysql *MySQL) InsertTodo(todo models.Todo) error {
	newTodo := models.Todo{
		Content:    todo.Content,
		IsFinished: todo.IsFinished,
		CreatedAt:  time.Now(),
	}

	mysql.db.Create(&newTodo)

	return nil
}

// GetAllTodos ...
func (mysql *MySQL) GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	mysql.db.Find(&todos)

	return todos, nil
}

// GetTodoWithID ...
func (mysql *MySQL) GetTodoWithID(id int) (models.Todo, error) {
	var todo models.Todo

	mysql.db.Where("id = ?", id).First(&todo)

	return todo, nil
}

// DeleteTodoWithID ...
func (mysql *MySQL) DeleteTodoWithID(id int) error {
	mysql.db.Where("id = ?", id).Delete(&models.Todo{})

	return nil
}

// DeleteAllTodos ...
func (mysql *MySQL) DeleteAllTodos() error {
	var todo models.Todo
	mysql.db.Delete(&todo)

	return nil
}
