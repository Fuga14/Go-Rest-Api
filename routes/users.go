package routes

import (
	"net/http"
	"rest-api/models"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse request data!",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not save user!",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse request data!",
		})
		return
	}

	err = user.ValidateLogin()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token!",
		})
		return
	}

	// Example of creation of cookie
	// context.SetCookie("token", token, 3600, "/", "localhost", false, true)

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
