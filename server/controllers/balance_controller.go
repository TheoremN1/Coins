package controllers

import (
	"net/http"

	"github.com/TheoremN1/Coins/database/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BalanceController struct {
	database *gorm.DB
}

func NewBalanceController(database *gorm.DB) *BalanceController {
	return &BalanceController{database}
}

func (balanceController *BalanceController) Get(context *gin.Context) {
	id := context.Request.URL.Query().Get("id")
	var status int
	var user models.User
	if id == "" {
		status = http.StatusBadRequest
	} else {
		var exists bool
		err := balanceController.database.
			Where("id = ?", id).
			Find(&exists).
			Error
		if err != nil || !exists {
			status = http.StatusBadRequest
		} else {
			status = http.StatusOK
			balanceController.database.
				Where("id = ?", id).
				First(user)
		}
	}

	context.JSON(status, gin.H{
		"balance": user.Balance,
		"status":  status,
	})
}
