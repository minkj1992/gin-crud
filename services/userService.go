package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	defaultCost = 10
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(origin string, given string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(origin), []byte[given])
	if err != nil {
		return false, fmt.Sprintf("password is incorrect")
	}
	return true, nil
}

