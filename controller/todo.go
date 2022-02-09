package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/entity"
	"github.com/minkj1992/gin-crud/service"
)

type TodoController interface {
	All()	[]entity.Todo
	Save(ctx *gin.Context) entity.Todo
}

type controller struct {
	service service.TodoService
}

func New(service service.TodoService) TodoController {
	return &controller{
		service: service,
	}
}


func (c *controller) All() []entity.Todo {
	return c.service.All()
}

func (c *controller) Save(ctx *gin.Context) entity.Todo {
	var todo entity.Todo
	ctx.BindJSON(&todo)
	c.service.Save(todo)
	return todo
}
