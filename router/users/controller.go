package users

import (
	"final-project/database"
	"final-project/router/middleware"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	userService := UserService{repository: UserRepository{db}}

	r.POST("/users/registerUser", userService.registerUser)
	r.GET("/users/loginUser", userService.loginUser)
	r.PUT("/users/updateUser", middleware.IsAuthorized(), userService.updateUser)
	r.DELETE("/users/deleteUser", middleware.IsAuthorized(), userService.deleteUser)
}
