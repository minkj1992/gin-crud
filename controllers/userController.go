package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/minkj1992/gin-crud/models"
	"github.com/minkj1992/gin-crud/services"
	"github.com/minkj1992/gin-crud/utils"
)

var validate = validator.New()

// debug only api
func GetUsers(c *gin.Context) {	
	var users []models.User
	_, cancel := utils.GetContextWithTimeOut()
	err := models.GetAllUsers(&users)
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, users)
}

// debug only api
func GetUser(c *gin.Context) {
	var user models.User
	
	userId := c.Param("user_id")
	err := models.GetUserById(&user, userId)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}

func GetCurrentUser(c * gin.Context) {
	var user models.User
	_, cancel := utils.GetContextWithTimeOut() // todo: do I need this line? wtf..

	uuid, ok := c.Get("uuid")
	if ok == false {
		cancel()
	}
	
	err := models.GetUserByUUID(&user, uuid.(string))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)	
}

func SignUp(c *gin.Context) {
	var user models.User
	_, cancel := utils.GetContextWithTimeOut() // todo: do I need this line? wtf..
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
	isExist, err := models.IsUserAlreadyExist(&user, *user.Email)
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking for the email"})
		return
	}
	if isExist {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "This user's email is already exists"})
		return
	}
	
	// hash password
	password := services.HashPassword(*user.Password)
	user.Password = &password	
	// generate uuid
	u := uuid.NewString()
	user.UUID = &u

	// TODO: uid만 전달 주면 될듯, register할 때는 토큰 만들지 않아도 될 것 같은데
	// 이렇게 되면 uuid도 필요없게 된다.

	//generate token and refersh token
	token, refreshToken, _ := services.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, *user.UUID)
	user.Token = &token
	user.RefreshToken = &refreshToken

	// register a user
	insertionError := models.CreateUser(&user)
	defer cancel()
	if insertionError != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	
	c.JSON(http.StatusOK, user.UUID)
}

func Login(c *gin.Context) {
	var user models.User
	var foundUser models.User
	// _, cancel := utils.GetContextWithTimeOut()

	//convert the login data from postman which is in JSON to golang readable format
	if err:=c.BindJSON(&user); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find a user with that email and see if that user even exists
	err := models.GetUserByEmail(&foundUser, *user.Email)
	// defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found, login seems to be incorrect"})
		return
	}
	
	// verify password
	ok, msg := services.VerifyPassword(*foundUser.Password, *user.Password)
	// defer cancel()
	if ok == false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	// update / generate tokens
	accessToken, refreshToken, _ := services.GenerateAllTokens(*foundUser.Email, *foundUser.FirstName, *foundUser.LastName, *foundUser.UUID)
	services.UpdateAllTokens(&foundUser, &accessToken, &refreshToken)
	c.JSON(http.StatusOK, accessToken)
}


func Withdraw(c *gin.Context) {
}