package photos

import (
	"final-project/database"
	"final-project/router/middleware"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	photoService := PhotoService{repository: PhotoRepository{db}}

	r.Use(middleware.IsAuthorized())
	r.POST("/photos/uploadPhoto", photoService.uploadPhoto)
	r.GET("/photos/getUserPhotos", photoService.getAllAssociateUserPhotos)
	// r.PUT("/users/updateUser", middleware.IsAuthorized(), userService.updateUser)
	// r.DELETE("/users/deleteUser", middleware.IsAuthorized(), userService.deleteUser)
}
