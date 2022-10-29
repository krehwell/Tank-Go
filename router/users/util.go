package users

import (
	"final-project/database"
	"final-project/model"
	"final-project/util"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func generateJWT(username, email string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["email"] = email
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(util.JWT_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}
