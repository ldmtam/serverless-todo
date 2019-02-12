package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ldmtam/serverless-todo/src/models"
	"github.com/ldmtam/serverless-todo/src/repositories"

	"github.com/gin-gonic/gin"
)

func getAllTodos(c *gin.Context) {
	db := c.MustGet("db").(repositories.Database)

	todoRepo := repositories.NewTodoRepo(db)

	todos, err := todoRepo.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get all todos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": todos,
	})
	return
}

func getTodoWithID(c *gin.Context) {
	db := c.MustGet("db").(repositories.Database)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "todo id must be integer",
		})
		return
	}

	todoRepo := repositories.NewTodoRepo(db)
	todo, err := todoRepo.GetTodoWithID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot get todo",
		})
		return
	}

	if todo.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("cannot find todo with id = %v", id),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": todo,
		})
		return
	}
}

func insertTodo(c *gin.Context) {
	db := c.MustGet("db").(repositories.Database)

	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot parse todo post data",
		})
		return
	}

	todoRepo := repositories.NewTodoRepo(db)
	err = todoRepo.InsertTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot insert todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "insert new todo successfully",
	})
	return
}

func deleteAllTodos(c *gin.Context) {
	db := c.MustGet("db").(repositories.Database)

	todoRepo := repositories.NewTodoRepo(db)
	err := todoRepo.DeleteAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot delete all todos",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "delete all todos succesfully",
	})
	return
}

func deleteTodoWithID(c *gin.Context) {
	db := c.MustGet("db").(repositories.Database)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "todo id must be integer",
		})
		return
	}

	todoRepo := repositories.NewTodoRepo(db)
	err = todoRepo.DeleteTodoWithID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("cannot delete todo with id = %v", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("delete todo with id = %v successfully", id),
	})
	return
}
