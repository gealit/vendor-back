package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	UserID uint64 `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
