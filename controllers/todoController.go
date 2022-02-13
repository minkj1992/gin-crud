package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/models"
)

func GetTodo(c *gin.Context) {	
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.GetTodoById(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodos(c *gin.Context) {
	var todo []models.Todo

	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	c.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.GetTodoById(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	// request body를 읽어 json.unmarshal하여 todo에 넣어준다. 이 과정에서 json 형태가 맞는지 validate한번 거친다.
	c.BindJSON(&todo) 
	err = models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}


func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.DeleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}