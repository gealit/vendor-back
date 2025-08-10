package controllers

import (
	"log"
	"main/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	accessToken, err := c.Cookie("access_token")

	if err != nil {
		log.Println("Cannot parse cookies...")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseAccessToken(accessToken)
	if err != nil {
		log.Println("Try to refresh tokens")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token_expires_in": claims.ExpiresAt,
		})
		return
	}

	newAccessToken, err := utils.RefreshTokens(c)
	if err != nil {
		log.Println("Cannot parse refresh token...")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
		"expires_in":   15 * 60, // 15 minutes in seconds
	})
}
