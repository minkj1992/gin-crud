package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/infra"
	"github.com/minkj1992/gin-crud/middleware"
	"github.com/minkj1992/gin-crud/models"
	"github.com/minkj1992/gin-crud/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	infra.LoadEnv()
	dsn := infra.Url(infra.Config())
	db, err := gorm.Open(mysql.Open(dsn),  &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to get db connection: %v", err)
	}

	infra.SetDB(db)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})

	router := gin.New()
	router.Use(gin.Logger())
	
	routes.BeforeAuthUserRoutes(router)
	router.Use(middleware.Authenticate())
	routes.AfterAuthUserRoutes(router)
	routes.TodoRoutes(router)
	router.Run(":8080")
}