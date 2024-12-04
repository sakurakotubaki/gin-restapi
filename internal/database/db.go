package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	// コネクションプールの設定
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance")
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
}

func GetDB() *gorm.DB {
	return DB
}

func AutoMigrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
