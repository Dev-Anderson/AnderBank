package routes

import (
	"github.com/dev-anderson/AnderBank/config"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	ConfigRoutes(router)

	port := config.LoadEnv().PortHttp
	router.Run("localhost:" + port)
}
