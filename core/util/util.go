package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("パスワードのハッシュかに失敗しました")
	}
	return string(hashedPassword), nil
}
