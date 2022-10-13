package router

import (
	"final-project/router/users"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	r := gin.Default()

	v1 := r.Group("/v1")

	users.InitializeUserRoutes(v1)

	r.Run(":8080")
}
