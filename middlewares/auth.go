package middlewares

import (
	"net/http"

	"example.com/go-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization") // get the token from the request header
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}
	userId, err := utils.VerifyToken(token) // validate the token
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized.", "error": err.Error()})
		return
	}
	context.Set("userId", userId) // set the userId in the context for use in the next handler
	context.Next()                // call the next handler in the chain

}
