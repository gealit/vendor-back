package middlewares

import (
	"log"
	"main/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			log.Println("Cannot parse cookies:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Парсим токен
		claims, err := utils.ParseAccessToken(accessToken)
		if err != nil {
			log.Println("Cannot parse token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Проверяем роль
		if claims.Role != "user" && claims.Role != "admin" {
			log.Println("Invalid role:", claims.Role)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Set("userEmail", claims.Email)
		c.Next()
	}
}
