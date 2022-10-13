package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
    r := gin.Default()

    v1 := r.Group("/v1")

    v1.GET("/home", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, "yakuza is cool")
    })

    r.Run(":8080")
}
