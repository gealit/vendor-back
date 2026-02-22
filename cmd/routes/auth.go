package routes

import (
	"main/cmd/controllers"
	"main/cmd/controllers/user"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/api/login", user.Login)
	r.POST("/api/signup", user.Signup)
	r.GET("/api", controllers.MainPage)
	r.GET("/api/me", user.GetMe)
	r.GET("/api/home", controllers.Home)
	r.GET("/api/premium", controllers.Premium)
	r.GET("/api/logout", user.Logout)
	r.GET("/api/refresh", controllers.RefreshToken)
}
