package photos

import (
	"final-project/database"
	"final-project/router/middleware"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	photoService := PhotoService{repository: PhotoRepository{db}}

	r.POST("/photos/uploadPhoto", middleware.IsAuthorized(), photoService.uploadPhoto)
	// r.GET("/users/loginUser", userService.loginUser)
	// r.PUT("/users/updateUser", middleware.IsAuthorized(), userService.updateUser)
	// r.DELETE("/users/deleteUser", middleware.IsAuthorized(), userService.deleteUser)
}
