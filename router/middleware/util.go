package middleware

import (
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

