package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(h), err
}

func CompareHashAndUserPassword(hashPass, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password)) == nil
}
