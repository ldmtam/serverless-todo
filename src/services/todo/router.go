package main

import (
	"net/http"

	"github.com/ldmtam/serverless-todo/src/repositories"
	"github.com/ldmtam/serverless-todo/src/utils"

	"github.com/gin-gonic/gin"
)

func dbMiddleware(mysql repositories.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", mysql)
		c.Next()
	}
}

func initRouter(mode string, mysql repositories.Database) *gin.Engine {
	if mode == utils.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(dbMiddleware(mysql))

	todoEndpoint := r.Group("/todo")
	{
		todoEndpoint.GET("", getAllTodos)
		todoEndpoint.GET("/:id", getTodoWithID)
		todoEndpoint.POST("", insertTodo)
		todoEndpoint.DELETE("", deleteAllTodos)
		todoEndpoint.DELETE("/:id", deleteTodoWithID)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,
			gin.H{
				"message": "API Not Found",
			},
		)
	})

	return r
}
