package utils

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) (uint64, bool) {
	value, exists := c.Get("userId")
	if !exists {
		return 0, false
	}

	user_id, ok := value.(uint64)
	return user_id, ok
}

func GetUserEmail(c *gin.Context) (string, bool) {
	value, exists := c.Get("userEmail")
	if !exists {
		return "", false
	}

	email, ok := value.(string)
	return email, ok
}

func GetUserRole(c *gin.Context) (string, bool) {
	value, exists := c.Get("userRole")
	if !exists {
		return "", false
	}

	role, ok := value.(string)
	return role, ok
}
