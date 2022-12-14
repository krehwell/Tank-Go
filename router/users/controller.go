package users

import (
	"final-project/database"
	"final-project/router/middleware"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	userService := UserService{repository: UserRepository{db}}

	r.POST("/users/register", userService.registerUser)
	r.GET("/users/login", userService.loginUser)
	r.PUT("/users/update", middleware.IsAuthorized(), userService.updateUser)
	r.DELETE("/users/delete", middleware.IsAuthorized(), userService.deleteUser)
}
