package models

import (
	"github.com/minkj1992/gin-crud/infra"
)


func GetAllUsers(users *[]User) (err error) {
	if err = infra.DB.Table("users").Find(users).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id string) (err error) {
	if err = infra.DB.Table("users").Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func IsUserAlreadyExist(user *User, email string) (bool, error) {
	var count int64
	if err := infra.DB.Table("users").Select("id").Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func CreateUser(user *User) (err error) {
	if err = infra.DB.Table("users").Create(user).Error; err !=nil {
		return err
	}
	return nil
}


func UpdateUser(user *User, id string) (err error) {
	if err = infra.DB.Table("users").Save(user).Error; err != nil {
		return err
	}
	return nil
	
}

func DeleteUser(user *User, id string) (err error) {
	if err = infra.DB.Table("users").Where("id = ?", id).Delete(user).Error; err !=nil {
		return err
	}
	return nil
}