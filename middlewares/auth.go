package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.naous.net/api/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "This is a protected route. Send token!"})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "This is a protected route. Authorize with token!", "error": err})
		return
	}

	context.Set("userid", userId)
	context.Next()

}
