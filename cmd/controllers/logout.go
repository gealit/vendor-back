package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", os.Getenv("HOST"), false, true)
	c.SetCookie("refresh_token", "", -1, "/", os.Getenv("HOST"), false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
