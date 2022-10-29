package utils

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	JWT_SECRET_KEY = "MY_JWT_SECRET_KEY"
)

func GenerateJWT(email, username string) (string, error) {
	var signingKey = []byte(JWT_SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsTokenValid(token string) bool {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return false
	}
	return true
}
