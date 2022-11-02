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

	if bindErr := ctx.ShouldBindJSON(&newUser); bindErr != nil {
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
		foundUser, isFound := u.repository.getUserByEmail(authData.Email)

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

	if bindErr := ctx.ShouldBindJSON(&user); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to bind with JSON, check the body"})
		return
	}

	oldUserData, jwtValidUserError := u.isUserIdEqualJwtUser(ctx, user.Id)
	if jwtValidUserError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": jwtValidUserError.Error()})
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

	if bindErr := ctx.ShouldBindJSON(&idToBeDeleted); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to bind with JSON, check the body"})
		return
	}

	userCorresId, jwtValidUserError := u.isUserIdEqualJwtUser(ctx, idToBeDeleted.Id)
	if jwtValidUserError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": jwtValidUserError.Error()})
		return
	}

	deletedUser, deleteUserErr := u.repository.updateUserData(userCorresId, model.User{IsDeleted: true})
	if deleteUserErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, deletedUser)
}
