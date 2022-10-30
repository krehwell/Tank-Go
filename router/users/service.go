package users

import (
	"errors"
	"final-project/model"
	"final-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repository UserRepository
}

func (u *UserService) registerUser(ctx *gin.Context) {
	newUser := model.User{}
	bindErr := ctx.ShouldBindJSON(&newUser)

	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": bindErr.Error()})
		return
	}

	createUser, createUserErr := u.repository.createNewUser(newUser)
	if createUserErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": createUserErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createUser)
}

func (u *UserService) loginUser(ctx *gin.Context) {
	authData := struct {
		Email    string
		Password string
	}{}

	if bindErr := ctx.ShouldBindJSON(&authData); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": bindErr.Error()})
		return
	}

	user, jwtToken, err := processUserAndGenerateToken(func() (model.User, error) {
		foundUser, isFound := u.repository.getUserByUsername(authData.Email)

		isAllowToLogin := utils.ComparePasswords(foundUser.Password, []byte(authData.Password))
		if !isAllowToLogin {
			return model.User{}, errors.New("User credential is invalid")
		}

		if !isFound {
			return foundUser, errors.New("User with given Id not found")
		}

		return foundUser, nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, struct {
		User  model.User
		Token string
	}{user, jwtToken})
}

func (u *UserService) updateUser(ctx *gin.Context) {
	user := model.User{}
	bindErr := ctx.ShouldBindJSON(&user)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to bind with JSON, check the body"})
		return
	}

	oldUserData, isFound := u.repository.getUserById(user.Id)
	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "User with given Id not found"})
		return
	}

	jwtUserNoAssertion, _ := ctx.Get(utils.JWT_USER_DATA_KEY)
	jwtUser := jwtUserNoAssertion.(utils.JWTUser)
	if oldUserData.Email != jwtUser.Email || oldUserData.Username != jwtUser.Username {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not authorized to update the user"})
		return
	}

	updatedUserData, jwtToken, err := processUserAndGenerateToken(func() (model.User, error) {
		return u.repository.updateUserData(oldUserData, user)
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, struct {
		User  model.User
		Token string
	}{updatedUserData, jwtToken})
}

func (u *UserService) deleteUser(ctx *gin.Context) {
	idToBeDeleted := struct{ Id string }{}

	bindErr := ctx.ShouldBindJSON(&idToBeDeleted)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to bind with JSON, check the body"})
		return
	}

	userCorresId, isFound := u.repository.getUserById(idToBeDeleted.Id)
	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "User with given Id not found"})
		return
	}

	jwtUserNoAssertion, _ := ctx.Get(utils.JWT_USER_DATA_KEY)
	jwtUser := jwtUserNoAssertion.(utils.JWTUser)
	if userCorresId.Email != jwtUser.Email || userCorresId.Username != jwtUser.Username {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Not authorized to delete the user"})
		return
	}

	deletedUser, deleteUserErr := u.repository.updateUserData(userCorresId, model.User{IsDeleted: 1})
	if deleteUserErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, deletedUser)
}
