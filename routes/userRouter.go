package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/controllers"
)

func BeforeAuthUserRoutes(incomingRoutes *gin.Engine) {
	users := incomingRoutes.Group("v1/users")
	{
		users.GET("/", controllers.GetUsers) // TODO:need api key
		users.GET("/:user_id", controllers.GetUser) // TODO:need api key
		users.POST("/signup", controllers.SignUp)
		users.POST("/login", controllers.Login)
	}
}


func AfterAuthUserRoutes(incomingRoutes *gin.Engine) {
	users := incomingRoutes.Group("v1/users")
	{
		// TODO: users.POST("/withdraw", controllers.Withdraw)
		// users.PATCH("/profile",controllers.GetCurrentUser)
		users.GET("/me",controllers.GetCurrentUser)
	}
}