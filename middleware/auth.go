package middleware

import (
	"net/http"
	"project/rest_api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized user"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err!=nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"not authourized"})
		return
	}

	context.Set("userId",userId)

	context.Next()
}