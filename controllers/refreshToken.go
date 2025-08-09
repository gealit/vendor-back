package controllers

import (
	"log"
	"main/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	_, err := c.Cookie("access_token")

	if err != nil {
		log.Println("Cannot parse cookies....")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	newAccessToken := utils.RefreshTokens(c)

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
		"expires_in":   15 * 60, // 15 minutes in seconds
	})
}
