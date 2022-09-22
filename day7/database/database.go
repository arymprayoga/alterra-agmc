package database

import (
	"fmt"
	"os"

	"day7/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var books = []model.Book{}

func CreateConnection() {

	config := map[string]string{
		"DB_Username": os.Getenv("DB_USERNAME"),
		"DB_Password": os.Getenv("DB_PASSWORD"),
		"DB_Port":     os.Getenv("DB_PORT"),
		"DB_Host":     os.Getenv("DB_HOST"),
		"DB_Name":     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.User{})
}

func GetConnection() *gorm.DB {
	if DB == nil {
		CreateConnection()
	}
	return DB
}
