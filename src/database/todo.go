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
