package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/models"
	"github.com/minkj1992/gin-crud/services"
	"github.com/minkj1992/gin-crud/utils"
)

var validate = validator.New()


func GetUser(c *gin.Context) {
	var user models.User
	
	userId := c.Param("user_id")
	err := models.GetUserById(&user, userId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func SignUp(c *gin.Context) {
	var user models.User
	_, cancel := utils.GetContextWithTimeOut()
	//convert the JSON data coming from postman to something that golang understands
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate the data based on user struct
	validationErr :=validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	// todo: Check user is already exist (unique email)
	// todo: hash password
	
	u := uuid.NewString()
	user.UUID = &u
	//generate token and refersh token (generate all tokens function from helper)
	token, refreshToken, _ := services.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UUID)
	user.Token = &token
	user.RefreshToken = &refreshToken

	insertionError := models.CreateUser(&user)
	defer cancel()
	if insertionError != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	
	//if all ok, then you insert this new user into the user collection	
	//return status OK and send the result back
	c.JSON(http.StatusOK, user.ID)
	
}
