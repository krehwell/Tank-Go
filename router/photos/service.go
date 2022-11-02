package photos

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PhotoService struct {
	repository PhotoRepository
}

func (p *PhotoService) uploadPhoto(ctx *gin.Context) {
	photoData := struct {
		Title    string
		Caption  string
		PhotoUrl string
	}{}

	bindErr := ctx.ShouldBindJSON(&photoData)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": bindErr.Error()})
		return
	}

}
