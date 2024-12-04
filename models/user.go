package models

import (
	"errors"
	"main/internal/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *User) Save() error {
	if u.ID != 0 {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return database.GetDB().Create(u).Error
}

func (u *User) ValidateCredentials() error {
	storedUser := &User{}
	db := database.GetDB()

	err := db.Where("email = ?", u.Email).First(storedUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	u.ID = storedUser.ID
	return nil
}
