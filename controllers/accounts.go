package controllers

import (
	"net/http"

	"github.com/dev-anderson/AnderBank/models"
	"github.com/gin-gonic/gin"
)

func GetAllAccounts(c *gin.Context) {
	accounts, err := models.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account failed" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func ConsultaTodos(c *gin.Context) {
	acc := models.Accounts{}
	accounts, err := acc.TodasContas()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account failed" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func CreateAccount(c *gin.Context) {
	var account models.Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON do Account invalido" + err.Error()})
		return
	}
	err = models.CreateAccount(account.Balance, account.Debit, account.Credit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nao foi possivel criar a conta" + err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func GetBalanceAccount(c *gin.Context) {
	number := c.Params.ByName("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nao foi informado o numero da conta, favor informar"})
	}
	balance, err := models.BalanceAccount(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falha ao carregar o saldo" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, balance)
}

func DeleteAccount(c *gin.Context) {
	number := c.Params.ByName("id")
	if number == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nao foi informado um numero de conta valido"})
		return
	}

	err := models.DeleteAccount(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nao foi possivel destaivar essa conta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Conta deleta com sucesso!"})
}
