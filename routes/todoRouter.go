package routes

import (
	controllers "github.com/minkj1992/gin-crud/controllers"

	"github.com/gin-gonic/gin"
)
func TodoRoutes(incomingRoutes *gin.Engine) {
	v1 := incomingRoutes.Group("/v1")
	v1.GET("todos", controllers.GetTodos)
	v1.POST("todos", controllers.CreateTodo)	
}