package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/ProductsService/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

type AchievementsController struct {
	databaseUrl string
}

func NewAchievementsController(databaseUrl string) *AchievementsController {
	return &AchievementsController{databaseUrl}
}

func (achievementsController *AchievementsController) Get(context *gin.Context) {
	url := achievementsController.databaseUrl + "/api/achievements"

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

	var achievements []models.Achievement
	json.NewDecoder(resp.Body).Decode(&achievements)

	context.JSON(resp.StatusCode, achievements)
}

func (achievementsController *AchievementsController) Post(context *gin.Context) {
	var achievement *models.Achievement
	json.NewDecoder(context.Request.Body).Decode(&achievement)

	url := achievementsController.databaseUrl + "/api/achievements"
	client := &http.Client{}
	req, err := sling.New().Post(url).BodyForm(achievement).Request()
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

func (achievementsController *AchievementsController) Put(context *gin.Context) {
	var achievement *models.Achievement
	json.NewDecoder(context.Request.Body).Decode(&achievement)

	url := achievementsController.databaseUrl + "/api/achievements/" + strconv.Itoa(achievement.Id)

	client := &http.Client{}
	req, err := sling.New().Put(url).BodyForm(achievement).Request()
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

func (achievementsController *AchievementsController) Delete(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := achievementsController.databaseUrl + "/api/achievements/" + id
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
