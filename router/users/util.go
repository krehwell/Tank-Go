package users

import (
	"final-project/database"
	"final-project/model"

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

func getUserByUsername(email string) (model.User, bool) {
	u := model.User{Email: email}
	if err := database.Instance.Db.First(&u, "email = ?", email).Error; err != nil {
		return u, false
	}

	return u, true
}
