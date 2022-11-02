package middleware

import (
	"final-project/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeader(c *gin.Context) string {
	headerStr := c.Request.Header.Get("Authorization")
	bearerTokenStr := strings.Split(headerStr, " ")

	if len(bearerTokenStr) != 2 {
		return ""
	}

	if bearerTokenStr[0] != "Bearer" {
		return ""
	}

	return bearerTokenStr[1]
}

func GetJWTUser(ctx *gin.Context) (utils.JWTUser, bool) {
	jwtUserPlain, isInContext := ctx.Get(utils.JWT_USER_DATA_KEY)
	if !isInContext {
		return utils.JWTUser{}, false
	}

	jwtUser, ok := jwtUserPlain.(utils.JWTUser)

	return jwtUser, ok
}
