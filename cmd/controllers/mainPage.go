package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, []gin.H{
		{
			"id":   100,
			"name": "Первый товар",
			"cost": 9999.99,
		},
		{
			"id":   101,
			"name": "Второй товар",
			"cost": 19999.99,
		},
	})
}
