package photos

import (
	"final-project/database"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	photoService := PhotoService{repository: PhotoRepository{db}}

	r.POST("/photos/uploadPhoto", photoService.uploadPhoto)
	// r.GET("/users/loginUser", userService.loginUser)
	// r.PUT("/users/updateUser", middleware.IsAuthorized(), userService.updateUser)
	// r.DELETE("/users/deleteUser", middleware.IsAuthorized(), userService.deleteUser)
}
