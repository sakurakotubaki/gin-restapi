package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string  `gorm:"unique;not null" binding:"required,email"`
	Password string  `gorm:"not null" binding:"required,min=6"`
	Events   []Event `gorm:"foreignKey:UserID"`
}

type Event struct {
	gorm.Model
	Name        string    `gorm:"not null" binding:"required"`
	Description string    `gorm:"not null" binding:"required"`
	Location    string    `gorm:"not null" binding:"required"`
	DateTime    time.Time `gorm:"not null" binding:"required"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
}
