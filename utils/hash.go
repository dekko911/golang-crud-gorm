package utils

import "golang.org/x/crypto/bcrypt"

func HashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashPassword(hashed string, plainTextPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), plainTextPassword)
	return err == nil
}
