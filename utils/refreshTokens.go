package utils

import (
	"fmt"
	"main/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func RefreshTokens(c *gin.Context) string {

	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token required"})
		return "error"
	}

	claims, err := ParseRefreshToken(refreshToken)
	if err != nil {
		fmt.Printf("Failed to parse refresh token.")
		return "error"
	}

	// Get user from DB
	var existingUser models.User

	models.DB.Where("id = ?", claims.UserID).First(&existingUser)

	if existingUser.ID == 0 {
		fmt.Printf("user does not exist with id: %v", claims.UserID)
		return "error"
	}

	// Generate new tokens
	newAccessToken, newRefreshToken, err := GenerateTokens(&existingUser)
	if err != nil {
		fmt.Printf("Failed to generate tokens")
		return "error"
	}

	// Set new cookies
	SetTokenCookies(c, newAccessToken, newRefreshToken)

	return newAccessToken
}

func SetTokenCookies(c *gin.Context, accessToken, refreshToken string) {
	// Secure, HttpOnly cookies
	c.SetSameSite(http.SameSiteStrictMode)

	// Access Token Cookie (short-lived)
	c.SetCookie(
		"access_token",
		accessToken,
		int((15 * time.Minute).Seconds()), // expires in 15 mins
		"/",
		os.Getenv("HOST"),
		false, // secure
		true,  // httpOnly
	)

	// Refresh Token Cookie (long-lived)
	c.SetCookie(
		"refresh_token",
		refreshToken,
		int((7 * 24 * time.Hour).Seconds()), // expires in 7 days
		"/",                                 // only accessible on refresh endpoint
		os.Getenv("HOST"),
		false, // secure
		true,  // httpOnly
	)
}
