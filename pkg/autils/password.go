package autils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashed string, password string) (bool, error) {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		return false, err
	}

	return true, err
}
