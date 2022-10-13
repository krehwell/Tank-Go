package users

import "github.com/gin-gonic/gin"

func InitializeUserRoutes(r *gin.RouterGroup) {
	r.POST("/users/register", registerUser)
	r.GET("/users/login", loginUser)
}
