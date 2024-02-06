package utils

import (
	"golang.org/x/crypto/bcrypt"

)

func HashPassword(password string) (string, error) {
	pass := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(hashPassword string, password string) error {
	pass := []byte(password)
	hashPass := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hashPass, pass)
	return err
}