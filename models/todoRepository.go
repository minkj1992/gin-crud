package models

import "github.com/minkj1992/gin-crud/infra"


func GetAllTodos(todo *[]Todo) (err error) {
	if err = infra.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err = infra.DB.Create(todo).Error; err !=nil {
		return err
	}
	return nil
}