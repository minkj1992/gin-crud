package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/controller"
	"github.com/minkj1992/gin-crud/service"
)


var (
	todoService service.TodoService = service.New()
	todoController controller.TodoController = controller.New(todoService)
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.GET("/todos", func(c *gin.Context){
		c.JSON(http.StatusOK, todoController.All())
	})
	server.POST("/todos", func(c *gin.Context){
		c.JSON(http.StatusCreated, todoController.Save(c))
	})
	
	server.Run(":8080")
}