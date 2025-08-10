package routes

import (
	"main/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/api/login", controllers.Login)
	r.POST("/api/signup", controllers.Signup)
	r.GET("/api", controllers.MainPage)
	r.GET("/api/home", controllers.Home)
	r.GET("/api/premium", controllers.Premium)
	r.GET("/api/logout", controllers.Logout)
	r.GET("/api/refresh", controllers.RefreshToken)
}
