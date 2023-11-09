package controllers

import (
	"net/http"

	"github.com/dev-anderson/AnderBank/models"
	"github.com/dev-anderson/AnderBank/schemas"
	"github.com/dev-anderson/AnderBank/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login schemas.Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nao foi possivel validar o JSON" + err.Error()})
		return
	}

	// var l schemas.Login
	user, err := models.Login(login.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao buscar o user"})
		return
	}

	if user.Password != services.Sha256Encoder(login.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha incorreta"})
		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o Token" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
}
