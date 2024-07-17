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
	// Input
	query := context.Request.URL.Query()
	id := query.Get("id")
	// Output
	var status int
	var hash gin.H

	if id != "" && balanceController.IsUserExist(id) {
		var user models.User
		balanceController.database.
			Where("id = ?", id).
			First(&user)
		status = http.StatusOK
		hash = gin.H{"balance": user.Balance}
	} else {
		status = http.StatusBadRequest
		hash = nil
	}

	context.JSON(status, hash)
}

func (balanceController *BalanceController) Put(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	id := query.Get("id")
	action := query.Get("action")
	amount, err := strconv.Atoi(query.Get("amount"))
	// Output
	var status int

	if (action == "plus" || action == "minus") &&
		id != "" &&
		err == nil &&
		amount > 0 &&
		balanceController.IsUserExist(id) {

		var user models.User
		balanceController.database.
			Where("id = ?", id).
			First(&user)

		switch action {
		case "plus":
			user.Balance += amount
		case "minus":
			user.Balance -= amount
		}

		balanceController.database.Save(&user)
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}

	context.JSON(status, nil)
}
