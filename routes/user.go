package routes

import (
	"fmt"

	"github.com/dsabljic/event-management/models"
	"github.com/dsabljic/event-management/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := user.Save()

	if err != nil {
		context.JSON(500, gin.H{"error": "Error creating an user", "message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	context.JSON(201, gin.H{"message": "user created successfully", "token": token})
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := user.ValidateCredentials()

	fmt.Println(user)

	if err != nil {
		context.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	context.JSON(200, gin.H{"message": "login successful", "token": token})
}
