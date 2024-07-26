package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/RequestsService/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

type MerchRequestController struct {
	databaseUrl string
}

func NewMerchRequestController(databaseUrl string) *MerchRequestController {
	return &MerchRequestController{databaseUrl}
}

func (merchRequestController *MerchRequestController) Get(context *gin.Context) {
	url := merchRequestController.databaseUrl + "/api/merchrequests"
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

	var requests []models.MerchRequest
	json.NewDecoder(resp.Body).Decode(&requests)

	context.JSON(resp.StatusCode, requests)
}

func (merchRequestController *MerchRequestController) Post(context *gin.Context) {
	var request *models.MerchRequest
	json.NewDecoder(context.Request.Body).Decode(&request)

	request.StatusKey = "wait"

	url := merchRequestController.databaseUrl + "/api/merchrequests"
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

func (merchRequestController *MerchRequestController) Put(context *gin.Context) {
	var request *models.MerchRequest
	json.NewDecoder(context.Request.Body).Decode(&request)

	url := merchRequestController.databaseUrl + "/api/merchrequests/" + strconv.Itoa(request.Id)

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

func (merchRequestController *MerchRequestController) Delete(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := merchRequestController.databaseUrl + "/api/merchrequests/" + id
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
