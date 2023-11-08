package controllers

import (
	"net/http"

	"github.com/dev-anderson/AnderBank/models"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	user, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	}
	c.JSON(http.StatusOK, user)
}
