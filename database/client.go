package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jwt-authentication-golang/models"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(conn string) {
	Instance, dbError = gorm.Open(mysql.Open(conn), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")
}
