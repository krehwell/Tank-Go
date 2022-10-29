package utils

import (
	"errors"
	"final-project/model"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	JWT_SECRET_KEY = "MY_JWT_SECRET_KEY"
	JWT_USER_DATA_KEY = "JWT_USER"
)

func parseToken(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(JWT_SECRET_KEY), nil
}

func ExtractTokenUserIdentity(tokenString string) (model.JWTUser, error) {
	token, parseErr := jwt.Parse(tokenString, parseToken)
	if parseErr != nil {
		return model.JWTUser{}, parseErr
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		username := claims["username"].(string)
		email := claims["email"].(string)

		return model.JWTUser{Username: username, Email: email}, nil
	}

	return model.JWTUser{}, errors.New("Token is invalid")
}

func GenerateJWT(email, username string) (string, error) {
	var signingKey = []byte(JWT_SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, parseErr := token.SignedString(signingKey)

	if parseErr != nil {
		return "", parseErr
	}
	return tokenString, nil
}

func IsTokenValid(token string) bool {
	_, err := jwt.Parse(token, parseToken)

	if err != nil {
		return false
	}
	return true
}
