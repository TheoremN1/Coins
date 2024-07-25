package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/UsersService/models"
	"github.com/dghubble/sling"
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
	client := &http.Client{}
	req, err := sling.New().Get(url).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user *models.User
	json.NewDecoder(resp.Body).Decode(&user)

	context.JSON(resp.StatusCode, user.Balance)
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
	client := &http.Client{}
	req, err := sling.New().Get(url).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user *models.User
	json.NewDecoder(resp.Body).Decode(&user)

	if user == nil {
		context.JSON(resp.StatusCode, false)
	} else {
		switch action {
		case "plus":
			user.Balance += amount
		case "minus":
			user.Balance -= amount
		default:
			panic("Unexpected action: " + action)
		}

		req, err := sling.New().Put(url).BodyForm(user).Request()
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		context.JSON(resp.StatusCode, true)
	}
}
