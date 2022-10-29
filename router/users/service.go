package users

import (
	"final-project/model"
	"final-project/utils"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserService struct {
	repository UserRepository
}

func (u *UserService) registerUser(ctx *gin.Context) {
	newUser := model.User{}
	bindErr := ctx.ShouldBindJSON(&newUser)

	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": bindErr.Error()})
		return
	}

	id := uuid.NewString()
	newUser.Id = id

	_, validateErr := govalidator.ValidateStruct(newUser)
	if validateErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": validateErr.Error()})
		return
	}
	hashedPassword := hashAndSalt([]byte(newUser.Password))
	newUser.Password = hashedPassword

	createUser, createUserErr := u.repository.createNewUser(newUser)
	if createUserErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": createUserErr.Error()})
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": bindErr.Error()})
		return
	}

	user, isFound := u.repository.getUserByUsername(authData.Email)
	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "User not found"})
		return
	}

	isAllowToLogin := comparePasswords(user.Password, []byte(authData.Password))
	if !isAllowToLogin {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "User credential Invalid"})
		return
	}

	jwtToken, jwtErr := utils.GenerateJWT(user.Email, user.Username)
	if jwtErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": jwtErr.Error})
	}

	ctx.JSON(http.StatusOK, struct {
		User  model.User
		Token string
	}{user, jwtToken})
}
