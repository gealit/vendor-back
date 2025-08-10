package controllers

import (
	"fmt"
	"main/cmd/utils"
	"main/internal/database"
	"main/internal/models"
	"main/internal/services"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	database.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	user.Role = "user"

	fmt.Println(user)

	s := services.NewUserService(database.DB)

	s.CreateUserWithProfile(&user, &models.Profile{})

	c.JSON(200, gin.H{"success": "user created"})
}
