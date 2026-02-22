package routes

import (
	"main/cmd/controllers"
	"main/cmd/controllers/user"
	"main/cmd/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/api/login", user.Login)
	r.POST("/api/signup", user.Signup)
	r.GET("/api", controllers.MainPage)
	r.GET("/api/refresh", controllers.RefreshToken)

	private := r.Group("/api")
	private.Use(middlewares.IsAuthorized())
	{
		private.GET("/me", user.GetMe)
		private.GET("/home", controllers.Home)
		private.GET("/premium", controllers.Premium)
		private.GET("/logout", user.Logout)
	}
}
