package infra

import (
	"fmt"

	"gorm.io/gorm"
)


var DB *gorm.DB


type DBConfig struct {
	Host		string
	Port		int
	User		string
	DBName 		string
	Password	string
}

func SetDB(db *gorm.DB) {
	DB = db
}

func Config() *DBConfig {
	password := GetEnv("DB_PASSWORD")
	return &DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "root",
		DBName:   "todos",
		Password: password,
	}
}

func Url(cfg *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}