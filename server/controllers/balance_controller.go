package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BalanceController struct{}

func NewBalanceController() *BalanceController {
	return &BalanceController{}
}

func (balanceController *BalanceController) Get(context *gin.Context) {
	id := context.GetInt("id")
	//TODO: запрос с бд о балансе
	context.JSON(http.StatusOK, gin.H{"id": id})
}
