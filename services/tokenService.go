package services

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/minkj1992/gin-crud/infra"
)

type SignedDetails struct {
	Email      string
	First_name string
	Last_name  string
	Uid        string
	jwt.StandardClaims
}

var secret_key string = infra.GetEnv("SECRET_KEY")
const (
	tokenTTL time.Duration = time.Hour * time.Duration(24)
	refreshTokenTTL time.Duration = time.Hour * time.Duration(24 * 7)
)


func createAccessToken(email string, firstName string, lastName string, uid string) (string, error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(tokenTTL).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
}

func createRefreshToken() (string, error) {
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(refreshTokenTTL).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))
}

func GenerateAllTokens(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {
	signedToken, err = createAccessToken(email, firstName, lastName, uid)
	signedRefreshToken, err = createRefreshToken()
	if err != nil {
		log.Panic(err)
		return
	}
	return signedToken, signedRefreshToken, err
}

func UpdateAllTokens() {
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret_key), nil
		},
	)
	
	//the token is invalid
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	//the token is expired
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprint("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}

