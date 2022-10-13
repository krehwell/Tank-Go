package users

import "github.com/gin-gonic/gin"

func getUsers(ctx *gin.Context) {
    ctx.JSON(200, "get users!")
}
