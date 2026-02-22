package user

import (
	"log"
	"main/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {
	// Получаем токен из куки
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		log.Println("Cannot parse cookies:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Парсим токен
	claims, err := utils.ParseAccessToken(accessToken)
	if err != nil {
		log.Println("Cannot parse token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Проверяем роль
	if claims.Role != "user" && claims.Role != "admin" {
		log.Println("Invalid role:", claims.Role)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Возвращаем данные пользователя
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    claims.UserID,
			"email": claims.Email,
			"role":  claims.Role,
			// добавьте другие поля, которые есть в claims
		},
	})
}
