package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(userID string, secretKey string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = userID
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
