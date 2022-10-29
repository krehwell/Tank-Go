package users

import (
	"final-project/model"
	"final-project/utils"

	"github.com/k0kubun/pp"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		pp.Println(err)
	}

	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		pp.Println(err)
		return false
	}

	return true
}

func processUserAndGenerateToken(cb func() (model.User, error)) (model.User, string, error) {
	user, processUserErr := cb()
	if processUserErr != nil {
		return model.User{}, "", processUserErr
	}

	jwtToken, jwtErr := utils.GenerateJWT(user.Email, user.Username)
	if jwtErr != nil {
		return model.User{}, "", jwtErr
	}

	return user, jwtToken, nil
}

