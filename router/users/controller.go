package users

import (
	"final-project/database"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	userService := UserService{repository: UserRepository{db}}

	r.POST("/users/register", userService.registerUser)
	r.GET("/users/login", userService.loginUser)
}
