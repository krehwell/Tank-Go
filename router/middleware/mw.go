package middleware

import (
	"final-project/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractToken(c *gin.Context) string {
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

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := extractToken(ctx)
		isTokenValid := utils.IsTokenValid(headerToken)

		if !isTokenValid {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Unauthorized"})
			return
		}
		ctx.Next()
	}
}
