package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required" gorm:"unique"`
	PasswordDigest string `json:"password_digest" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}
