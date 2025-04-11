package middlewares

import (
	"net/http"
	"strings"

	"example.com/event/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorize"})
		return
	}

	var bearertoken = "Bearer "

	ok := strings.Contains(token, bearertoken)

	if ok {
		token = token[len(bearertoken):]
	}

	userID, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorize", "error": err.Error()})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
