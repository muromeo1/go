package auth

import (
	"errors"
	"go-api/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func SessionCreator(login Login) (string, error) {
	db := config.PG()
	user := User{}
	rows := db.Where("email = ?", login.Email).Find(&user).RowsAffected

	if rows == 0 {
		err := errors.New("user not found")
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(login.Password)); err != nil {
		err := errors.New("invalid password")
		return "", err
	}

	token, err := TokenCreator(user)

	if err != nil {
		err := errors.New("failed to create token")
		return "", err
	}

	return token, nil
}
