package service

import "github.com/minkj1992/gin-crud/entity"

type TodoService interface {
	Save(entity.Todo)	entity.Todo
	All()				[]entity.Todo
}


type todoService struct {
	todoList []entity.Todo
}


func New() TodoService {
	return &todoService{
		todoList: []entity.Todo{},
	}
}

func (service *todoService) Save(todo entity.Todo) entity.Todo {
	service.todoList = append(service.todoList, todo)
	return todo
}
func (service *todoService) All() []entity.Todo {
	return service.todoList

}