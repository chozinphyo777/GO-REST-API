package routes

import (
	"net/http"

	"example.com/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}
	err = user.Save() // save the user
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
	}
	err = user.ValidateCredentials() // validate the user's credentials
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful"})

}
