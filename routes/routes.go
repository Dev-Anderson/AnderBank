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
			user.DELETE("/delete/:id", controllers.DeleteUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		account := main.Group("account")
		{
			account.GET("/", controllers.GetAllAccounts)
			account.GET("/todos/", controllers.ConsultaTodos)
			account.POST("/", controllers.CreateAccount)
			account.GET("/balance/:id", controllers.GetBalanceAccount)
			account.POST("/delete/:id", controllers.DeleteAccount)
		}
	}

	return router
}
