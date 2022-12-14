package photos

import (
	"final-project/database"
	"final-project/router/middleware"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	photoService := PhotoService{repository: PhotoRepository{db}}

	r.Use(middleware.IsAuthorized())
	r.POST("/photos/upload", photoService.uploadPhoto)
	r.GET("/photos/getAll", photoService.getAllAssociateUserPhotos)
	r.PUT("/photos/update", photoService.updatePhoto)
	r.DELETE("/photos/delete", middleware.IsAuthorized(), photoService.deletePhoto)
}
