package controllers

import (
	"net/http"
	"strconv"

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

	user.Password = services.Sha256Encoder(user.Password)

	err = models.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nao foi possivel criar o usuario" + err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func GetIDUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID usuario invalido"})
		return
	}

	user, err := models.GetUserID(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao buscar o user por ID"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUserDelete(c *gin.Context) {
	user, err := models.GetAllUserDelete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao buscar usuarios deletados"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID usuario invalido"})
		return
	}

	err = models.DeleteIDUser(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao deletar user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario deletado com sucesso"})
}
