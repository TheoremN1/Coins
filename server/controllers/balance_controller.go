package controllers

import (
	"net/http"
	"strconv"

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

func (balanceController *BalanceController) IsUserExist(id string) bool {
	var users []models.User
	balanceController.database.
		Where("id = ?", id).
		Find(&users)
	return len(users) == 1
}

func (balanceController *BalanceController) Get(context *gin.Context) {
	id := context.Request.URL.Query().Get("id")
	var status int
	var user models.User
	if id == "" {
		status = http.StatusBadRequest
	} else {
		if balanceController.IsUserExist(id) {
			status = http.StatusOK
			balanceController.database.
				Where("id = ?", id).
				First(&user)
		} else {
			status = http.StatusBadRequest
		}
	}

	context.JSON(status, gin.H{
		"balance": user.Balance,
		"status":  status,
	})
}

func (balanceController *BalanceController) Put(context *gin.Context) {
	query := context.Request.URL.Query()
	var status int
	id := query.Get("id")
	if id != "" && balanceController.IsUserExist(id) {
		var user models.User
		balanceController.database.
			Where("id = ?", id).
			First(&user)
		switch query.Get("action") {
		case "plus":
			amount, err := strconv.Atoi(query.Get("amount"))
			if err == nil {
				user.Balance += amount
				balanceController.database.Save(&user)
				status = http.StatusOK
			} else {
				status = http.StatusBadRequest
			}
		case "minus":
			amount, err := strconv.Atoi(query.Get("amount"))
			if err == nil {
				user.Balance -= amount
				balanceController.database.Save(&user)
				status = http.StatusOK
			} else {
				status = http.StatusBadRequest
			}
		default:
			status = http.StatusBadRequest
		}
	} else {
		status = http.StatusBadRequest
	}

	context.JSON(status, gin.H{})
}
