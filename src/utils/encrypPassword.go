package utils

import "golang.org/x/crypto/bcrypt"

func EncrypPassword(password string) (string, error) {
	cost := Config.Costencrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
