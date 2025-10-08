package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/what-crud/initializers"
)

func CreateJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(3 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(initializers.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
