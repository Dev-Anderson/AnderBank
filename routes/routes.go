package routes

import (
	"github.com/dev-anderson/AnderBank/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUser)
			user.POST("/", controllers.CreateUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
	}

	return router
}
