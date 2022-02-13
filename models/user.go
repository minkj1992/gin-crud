package models

import (
	"time"
)

// todo: BASE model: https://stackoverflow.com/questions/36486511/how-do-you-do-uuid-in-golangs-gorm
type User struct {
	ID      		uint   		`json:"id"`
	UUID 			*string 	`json:"uuid"`
	FirstName		*string 	`json:"first_name" validate:"required,min=2,max=100"`
	LastName		*string 	`json:"last_name" validate:"required,min=2,max=100"`
	Email 			*string 	`json:"email" validate:"email,required"`
	Password		*string 	`json:"password" validate:"required,min=6"`
	Avatar 			*string 	`json:"avatar"`
	Token			*string 	`json:"token"`
	RefreshToken	*string 	`json:"refresh_token"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}

func (_ *User) TableName() string {
	return "users"
}