package users

import (
	"final-project/model"
	"final-project/utils"
)

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

