package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type BalanceController struct {
	databaseUrl string
}

func NewBalanceController(databaseUrl string) *BalanceController {
	return &BalanceController{databaseUrl}
}

func (balanceController *BalanceController) Get(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := balanceController.databaseUrl + "/api/users/" + id

	user, status := GetUser(url)

	context.JSON(status, user.Balance)
}

func (balanceController *BalanceController) Put(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")
	action := query.Get("action")
	amount, err := strconv.Atoi(query.Get("amount"))
	if err != nil {
		panic(err)
	}

	url := balanceController.databaseUrl + "/api/users/" + id

	user, status := GetUser(url)

	if user == nil {
		context.JSON(status, false)
	} else {
		switch action {
		case "plus":
			user.Balance += amount
		case "minus":
			user.Balance -= amount
		default:
			panic("Unexpected action: " + action)
		}

		isPuted, status := EditUser(url, user)

		context.JSON(status, isPuted)
	}
}
