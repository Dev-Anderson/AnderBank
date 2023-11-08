package middleware

import (
	"net/http"

	"github.com/dev-anderson/AnderBank/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
