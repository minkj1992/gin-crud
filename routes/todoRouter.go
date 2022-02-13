package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/controllers"
)

func TodoRoutes(incomingRoutes *gin.Engine) {
	todos := incomingRoutes.Group("v1/todos")
	{
		todos.GET("/:id", controllers.GetTodo)
		todos.GET("/", controllers.GetTodos)
		todos.POST("/", controllers.CreateTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}
}