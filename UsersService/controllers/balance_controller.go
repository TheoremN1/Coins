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
	context.JSON(http.StatusOK, gin.H{"type": "get"})
}

func (balanceController *BalanceController) Put(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"type": "put"})
}
