package users

import (
	"final-project/database"
	"final-project/model"
	"final-project/util"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func registerUser(ctx *gin.Context) {
	newUser := model.User{}
	bindErr := ctx.ShouldBindJSON(&newUser)

	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: bindErr.Error()})
		return
	}

	id := uuid.NewString()
	newUser.Id = id

	_, validateErr := govalidator.ValidateStruct(newUser)
	if validateErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: validateErr.Error()})
		return
	}
	hashedPassword := hashAndSalt([]byte(newUser.Password))
	newUser.Password = hashedPassword

	createUserErr := database.Instance.Db.Create(&newUser).Error
	if createUserErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: createUserErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}

func loginUser(ctx *gin.Context) {
	authData := struct {
		Email    string
		Password string
	}{}

	bindErr := ctx.ShouldBindJSON(&authData)

	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: bindErr.Error()})
		return
	}

	// is user exist
	user, isFound := getUserByUsername(authData.Email)
	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: "User not found!"})
		return
	}

	// is password user correct
	isAllowToLogin := comparePasswords(user.Password, []byte(authData.Password))
	if !isAllowToLogin {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: "User credential is invalid"})
		return
	}

	loginToken, tokenErr := generateJWT(user.Username, user.Email)
	if tokenErr != nil {
		fmt.Println("token error", tokenErr)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: "Error in generating token for user"})
	}

	ctx.JSON(http.StatusOK, loginToken)
}
