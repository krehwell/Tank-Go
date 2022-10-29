package router

import (
	"final-project/database"
	"final-project/router/users"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(db database.Database) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	users.Controller(v1, db)

	return r
}
