package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func (user *User) PasswordHasher(p string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	if err != nil {
		return err
	}
	user.Password = string(b)
	return nil
}

func (user *User) PasswordChecker(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
