package middlewares

import (
	"fmt"
	"main/cmd/utils"
	"main/internal/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		access, _ := c.Cookie("access_token")
		refresh, _ := c.Cookie("refresh_token")
		fmt.Println("access_token: ", access, "refresh_token: ", refresh)

		if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/refresh" {
			c.Next()
			return
		}

		secretKey := []byte(os.Getenv("SECRET_KEY"))
		// Try to get access token
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			oldClaims, _ := token.Claims.(*models.Claims)
			// If access token is expired, try to refresh it
			if oldClaims.ExpiresAt < time.Now().Unix() {
				utils.RefreshTokens(c)
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Set("userEmail", claims.Email)
		c.Next()
	}
}
