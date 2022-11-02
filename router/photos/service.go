package photos

import (
	"final-project/model"
	"final-project/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoService struct {
	repository PhotoRepository
}

func (p *PhotoService) uploadPhoto(ctx *gin.Context) {
	photoData := model.Photo{}
	if bindErr := ctx.ShouldBindJSON(&photoData); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": bindErr.Error()})
		return
	}

	jwtUser, ok := middleware.GetJWTUser(ctx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "No user in found in JWT"})
		return
	}

	photoData.UserId = jwtUser.Id

	createdPhoto, createPhotoErr := p.repository.createPhoto(photoData)
	if createPhotoErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": createPhotoErr.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdPhoto)
}

func (p *PhotoService) getAllAssociateUserPhotos(ctx *gin.Context) {

}
