package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		return
	}

	userId, err := VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalid!",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()

}
