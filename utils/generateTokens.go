package utils

import (
	"main/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateTokens(user *models.User) (accessToken string, refreshToken string, err error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	// Access Token (short-lived, 15-30 minutes)
	accessExpirationTime := time.Now().Add(2 * time.Minute)
	accessClaims := &models.Claims{
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token (long-lived, 7-30 days)
	refreshExpirationTime := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &models.RefreshClaims{
		Email:  user.Email,
		UserID: uint64(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)

	return accessToken, refreshToken, err
}
