package controllers

import (
	"fmt"
	"net/http"

	"github.com/dev-anderson/AnderBank/models"
	"github.com/dev-anderson/AnderBank/services"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	user, err := models.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nao foi possivel ler o JSON" + err.Error(),
		})
		return
	}

	fmt.Println(user)

	user.Password = services.Sha256Encoder(user.Password)

	err = models.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nao foi possivel criar o usuario" + err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
