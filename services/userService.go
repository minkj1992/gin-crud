package services

import (
	"fmt"
	"log"

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

func VerifyPassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, fmt.Sprintf("password is incorrect")
	}
	return true, ""
}