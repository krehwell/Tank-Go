package users

import (
	"errors"
	"final-project/model"
	"final-project/utils"

	"final-project/router/middleware"

	"github.com/gin-gonic/gin"
)

func processUserAndGenerateToken(cb func() (model.User, error)) (model.User, string, error) {
	user, processUserErr := cb()
	if processUserErr != nil {
		return model.User{}, "", processUserErr
	}

	jwtToken, jwtErr := utils.GenerateJWT(user.Email, user.Username, user.Id)
	if jwtErr != nil {
		return model.User{}, "", jwtErr
	}

	return user, jwtToken, nil
}

func (u *UserService) isUserIdEqualJwtUser(ctx *gin.Context, userId string) (model.User, error) {
	userCorresId, isFound := u.repository.getUserById(userId)
	if !isFound {
		return model.User{}, errors.New("User with given Id not found")
	}

	jwtUser, ok := middleware.GetJWTUser(ctx)
	if userCorresId.Email != jwtUser.Email || userCorresId.Username != jwtUser.Username  || !ok {
		return model.User{}, errors.New("User in JWT missmatch")
	}

	return userCorresId, nil
}
