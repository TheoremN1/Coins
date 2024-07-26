package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/ProductsService/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

type MerchController struct {
	databaseUrl string
}

func NewMerchController(databaseUrl string) *MerchController {
	return &MerchController{databaseUrl}
}

func (merchController *MerchController) Get(context *gin.Context) {
	url := merchController.databaseUrl + "/api/merch"

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

	var merch []models.Merch
	json.NewDecoder(resp.Body).Decode(&merch)

	context.JSON(resp.StatusCode, merch)
}

func (merchController *MerchController) Post(context *gin.Context) {
	var merch *models.Merch
	json.NewDecoder(context.Request.Body).Decode(&merch)

	url := merchController.databaseUrl + "/api/merch"
	client := &http.Client{}
	req, err := sling.New().Post(url).BodyForm(merch).Request()
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

func (merchController *MerchController) Put(context *gin.Context) {
	var merch *models.Merch
	json.NewDecoder(context.Request.Body).Decode(&merch)

	url := merchController.databaseUrl + "/api/merch/" + strconv.Itoa(merch.Id)

	client := &http.Client{}
	req, err := sling.New().Put(url).BodyForm(merch).Request()
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

func (merchController *MerchController) Delete(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := merchController.databaseUrl + "/api/merch/" + id
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
