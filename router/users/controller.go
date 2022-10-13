package users

import (
	"final-project/database"
	"final-project/model"
	"final-project/util"
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

	createUserErr := database.Instance.Db.Create(&newUser).Error
	if createUserErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.ErrMsg{ErrorMessage: validateErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newUser)
}
