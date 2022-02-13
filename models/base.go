package models

import (
	"time"

	"github.com/minkj1992/gin-crud/utils"
	"gorm.io/gorm"
)

type Base struct {
	ID      	uint   `json:"id"`
	CreatedAt 		time.Time 	`json:"created_at"`
	UpdatedAt 		time.Time 	`json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	currentDateTime, err := utils.CurrentDateTime()
	if err != nil {
	  return err
	}
	tx.Statement.SetColumn("CreatedAt", currentDateTime)
	tx.Statement.SetColumn("UpdatedAt", currentDateTime)
	return nil
}

