package controllers

import (
	"main/cmd/utils"
	"main/internal/database"
	"main/internal/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	database.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(&existingUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate tokens"})
		return
	}

	utils.SetTokenCookies(c, accessToken, refreshToken)

	c.JSON(200, gin.H{"message": "Logged in successfully", "user": user.Email})
}
