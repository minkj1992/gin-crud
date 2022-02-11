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

func GetTodoById(todo *Todo, id string) (err error) {
	if err = infra.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo, id string) (err error) {
	if err = infra.DB.Save(todo).Error; err != nil {
		return err
	}
	return nil
	
}

func DeleteTodo(todo *Todo, id string) (err error) {
	if err = infra.DB.Where("id = ?", id).Delete(todo).Error; err !=nil {
		return err
	}
	return nil
}