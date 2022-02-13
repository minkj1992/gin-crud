package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	users := incomingRoutes.Group("v1/users")
	{
		// users.GET("/", controllers.GetUsers)
		// users.GET("/:user_id", controllers.GetUser)
		users.POST("/signup", controllers.SignUp)
		users.POST("/login", controllers.Login)
		users.POST("/logout", controllers.Logout)
	}
}