package db

import (
	"log"
	"main/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("events.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	// Auto Migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Event{})
	if err != nil {
		log.Fatal("Failed to migrate database")
	}
}
