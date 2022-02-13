package models

type User struct {
	Base
	UUID 			*string 	`json:"uuid" gorm:"index:idx_uuid,unique"`
	FirstName		*string 	`json:"first_name" validate:"required,min=2,max=100"`
	LastName		*string 	`json:"last_name" validate:"required,min=2,max=100"`
	Email 			*string 	`json:"email" validate:"email,required" gorm:"index:idx_email,unique"`
	Password		*string 	`json:"password" validate:"required,min=6"`
	Avatar 			*string 	`json:"avatar"`
	Token			*string 	`json:"token"`
	RefreshToken	*string 	`json:"refresh_token"`
}

func (_ *User) TableName() string {
	return "users"
}