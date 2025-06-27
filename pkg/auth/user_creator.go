package auth

import (
	"errors"
	"go-api/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password *string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	return string(hashed)
}

func UserCreator(register Register) (string, error) {
	if register.Password != register.PasswordConfirmation {
		err := errors.New("password and confirmation do not match")
		return "", err
	}

	user := User{
		Name:           register.Name,
		Email:          register.Email,
		PasswordDigest: hashPassword(&register.Password),
	}

	db := config.PG()
	rows := db.Where("email = ?", user.Email).Find(&user).RowsAffected

	if rows != 0 {
		err := errors.New("user already exists")
		return "", err
	}

	db.Create(&user)

	token, _ := TokenEncode(user)

	return token, nil
}
