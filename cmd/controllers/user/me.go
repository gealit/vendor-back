package user

import (
	"main/cmd/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMe(c *gin.Context) {

	user_id, _ := utils.GetUserID(c)
	user_email, _ := utils.GetUserEmail(c)
	user_role, _ := utils.GetUserRole(c)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":    user_id,
			"email": user_email,
			"role":  user_role,
		},
	})
}
