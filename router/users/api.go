package users

import "github.com/gin-gonic/gin"

func InitializeUserRoutes(r *gin.RouterGroup) {
	r.GET("/users", getUsers)
}
