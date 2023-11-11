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
			user.GET("/:id", controllers.GetIDUser)
			user.GET("/delete/", controllers.GetAllUserDelete)
			user.POST("/", controllers.CreateUser)
			user.POST("/delete/:id", controllers.DeleteUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
	}

	return router
}
