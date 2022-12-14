package photos

import (
	"final-project/model"
	"final-project/router/middleware"
	"final-project/utils"
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
	jwtUser, ok := middleware.GetJWTUser(ctx)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "No user in found in JWT"})
		return
	}

	photos, getAllPhotosErr := p.repository.getPhotosByUserId(jwtUser.Id)

	if getAllPhotosErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": getAllPhotosErr.Error()})
		return
	}

	type PhotoResponse struct {
		model.Photo
		User utils.JWTUser
	}

	result := []PhotoResponse{}
	for i := range photos {
		result = append(result, PhotoResponse{photos[i], jwtUser})
	}

	ctx.JSON(http.StatusOK, result)
}

func (p *PhotoService) updatePhoto(ctx *gin.Context) {
	photoBody := model.PhotoBody{}
	if bindErr := ctx.ShouldBindJSON(&photoBody); bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": bindErr.Error()})
		return
	}

	oldPhotoData, getPhotoErr := p.repository.getPhotoById(photoBody.Id)
	if getPhotoErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": getPhotoErr.Error()})
		return
	}

	newPhotoData := oldPhotoData
	utils.MergeInPlaceStructWithPartialStruct(&newPhotoData, photoBody)

	updatedPhoto, updatePhotoErr := p.repository.updatePhotoData(oldPhotoData, newPhotoData)
	if updatePhotoErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": updatePhotoErr.Error()})
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}
