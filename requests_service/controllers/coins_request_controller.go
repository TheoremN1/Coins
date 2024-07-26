package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/RequestsService/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

type CoinsRequestController struct {
	databaseUrl string
}

func NewCoinsRequestController(databaseUrl string) *CoinsRequestController {
	return &CoinsRequestController{databaseUrl}
}

func (coinsRequestController *CoinsRequestController) Get(context *gin.Context) {
	url := coinsRequestController.databaseUrl + "/api/coinsrequests"
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

	var requests []models.CoinsRequest
	json.NewDecoder(resp.Body).Decode(&requests)

	context.JSON(resp.StatusCode, requests)
}

func (coinsRequestController *CoinsRequestController) Post(context *gin.Context) {
	var request *models.CoinsRequest
	json.NewDecoder(context.Request.Body).Decode(&request)

	request.StatusKey = "wait"

	url := coinsRequestController.databaseUrl + "/api/coinsrequests"
	client := &http.Client{}
	req, err := sling.New().Post(url).BodyForm(request).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var isPosted bool
	json.NewDecoder(resp.Body).Decode(&isPosted)

	context.JSON(resp.StatusCode, isPosted)
}

func (coinsRequestController *CoinsRequestController) Put(context *gin.Context) {
	var request *models.CoinsRequest
	json.NewDecoder(context.Request.Body).Decode(&request)

	url := coinsRequestController.databaseUrl + "/api/coinsrequests/" + strconv.Itoa(request.Id)

	client := &http.Client{}
	req, err := sling.New().Put(url).BodyForm(request).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var isPuted bool
	json.NewDecoder(resp.Body).Decode(&isPuted)

	context.JSON(resp.StatusCode, isPuted)
}

func (coinsRequestController *CoinsRequestController) Delete(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := coinsRequestController.databaseUrl + "/api/coinsrequests/" + id
	client := &http.Client{}
	req, err := sling.New().Delete(url).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var isDeleted *bool
	json.NewDecoder(resp.Body).Decode(&isDeleted)

	context.JSON(resp.StatusCode, isDeleted)
}
