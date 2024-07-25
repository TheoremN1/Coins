package controllers

import "github.com/gin-gonic/gin"

type AchievementsController struct {
	databaseUrl string
}

func NewAchievementsController(databaseUrl string) *AchievementsController {
	return &AchievementsController{databaseUrl}
}

func (achievementsController *AchievementsController) Get(context *gin.Context) {

}

func (achievementsController *AchievementsController) Post(context *gin.Context) {

}

func (achievementsController *AchievementsController) Put(context *gin.Context) {

}

func (achievementsController *AchievementsController) Delete(context *gin.Context) {

}
