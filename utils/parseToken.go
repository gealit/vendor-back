package utils

import (
	"fmt"
	"log"
	"main/models"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func ParseAccessToken(tokenString string) (*models.Claims, error) {

	secretKey := []byte(os.Getenv("SECRET_KEY"))

	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("Failed to parse token: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		fmt.Printf("UserRole: %v, Email: %v\n", claims.Role, claims.Email)
		fmt.Printf("Expires at: %v\n", claims.ExpiresAt)
	} else {
		log.Fatal("Invalid token")
	}

	return claims, nil
}

func ParseRefreshToken(tokenString string) (*models.RefreshClaims, error) {

	secretKey := []byte(os.Getenv("SECRET_KEY"))

	claims := &models.RefreshClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("Failed to parse token: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*models.RefreshClaims); ok && token.Valid {
		fmt.Printf("UserID: %v, Email: %v\n", claims.UserID, claims.Email)
		fmt.Printf("Expires at: %v\n", claims.ExpiresAt)
	} else {
		log.Fatal("Invalid token")
	}

	return claims, nil
}
