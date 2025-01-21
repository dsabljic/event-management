package middleware

import (
	"github.com/dsabljic/event-management/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(401, gin.H{"error": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(401, gin.H{"error": "Not authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
