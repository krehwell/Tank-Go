package utils

import (
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
