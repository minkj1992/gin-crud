package infra

import (
	"fmt"
	"log"
	"strconv"

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
	var port int
	var err error
	
	if port, err = strconv.Atoi(GetEnv("DB_PORT")); err != nil {
        log.Fatalf("Error loading .env file")
    } 
	
	return &DBConfig{
		Host:     GetEnv("DB_HOST"),
		Port:     port,
		User:     GetEnv("DB_USER"),
		DBName:   GetEnv("DB_NAME"),
		Password: GetEnv("DB_PASSWORD"),
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