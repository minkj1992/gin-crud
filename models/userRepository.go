package models

import "github.com/minkj1992/gin-crud/infra"


func GetAllUsers(users *[]User) (err error) {
	if err = infra.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = infra.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(user *User, email string) (err error) {
	if err = infra.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = infra.DB.Create(user).Error; err !=nil {
		return err
	}
	return nil
}


func UpdateUser(user *User, id string) (err error) {
	if err = infra.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
	
}

func DeleteUser(user *User, id string) (err error) {
	if err = infra.DB.Where("id = ?", id).Delete(user).Error; err !=nil {
		return err
	}
	return nil
}