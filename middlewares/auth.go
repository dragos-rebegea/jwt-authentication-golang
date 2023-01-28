package middlewares

import (
	"strings"

	"github.com/dragos-rebegea/jwt-authentication-golang/auth"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := strings.Split(context.Request.Header.Get("Authorization"), "Bearer ")
		if len(tokenString) != 2 {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString[1])
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
