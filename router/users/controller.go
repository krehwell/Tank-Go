package users

import (
	"final-project/database"

	"github.com/gin-gonic/gin"
)

func Controller(r *gin.RouterGroup, db database.Database) {
	userRepo := UserRepository{database: db}

	r.POST("/users/register", userRepo.registerUser)
	// r.GET("/users/login", loginUser)
}
