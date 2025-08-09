package controllers

import (
	"log"
	"main/utils"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	accessToken, err := c.Cookie("access_token")

	if err != nil {
		log.Println("Cannot parse cookies....")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseAccessToken(accessToken)

	if err != nil {
		log.Println("Cannot parse token....")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		log.Println("Cannot read the role....")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}
