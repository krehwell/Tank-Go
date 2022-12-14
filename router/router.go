package router

import (
	"final-project/database"
	"final-project/router/photos"
	"final-project/router/users"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(db database.Database) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	v1 := r.Group("/v1")

	users.Controller(v1, db)
	photos.Controller(v1, db)

	return r
}
