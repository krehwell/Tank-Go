package users

import (
	"errors"
	"final-project/model"
	"final-project/utils"

	"github.com/gin-gonic/gin"
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

func (u *UserService) isUserIdEqualJwtUser(ctx *gin.Context, userId string) (model.User, error) {
	userCorresId, isFound := u.repository.getUserById(userId)
	if !isFound {
		return model.User{}, errors.New("User with given Id not found")
	}

	jwtUserNoAssertion, _ := ctx.Get(utils.JWT_USER_DATA_KEY)
	jwtUser := jwtUserNoAssertion.(utils.JWTUser)
	if userCorresId.Email != jwtUser.Email || userCorresId.Username != jwtUser.Username {
		return model.User{}, errors.New("Not authorized to delete the user")
	}

	return userCorresId, nil
}
