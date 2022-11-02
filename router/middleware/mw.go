package middleware

import (
	"final-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerToken := extractTokenFromHeader(ctx)

		isTokenValid := utils.IsTokenValid(headerToken)

		if !isTokenValid {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Unauthorized"})
			return
		}

		jwtUserData, jwtUserDataErr := utils.ExtractTokenUserIdentity(headerToken)
		if jwtUserDataErr != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": jwtUserDataErr.Error()})
			return
		}

		ctx.Set(utils.JWT_USER_DATA_KEY, jwtUserData)
		ctx.Next()
	}
}
