package controllers

import (
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/database/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AchievementsController struct {
	database *gorm.DB
}

func NewAchievementsController(database *gorm.DB) *AchievementsController {
	return &AchievementsController{database}
}

func (achievementsController *AchievementsController) Get(context *gin.Context) {
	// Output
	var status int
	var hash gin.H

	var achievements []models.Achievement
	achievementsController.database.
		Find(&achievements)

	if len(achievements) > 0 {
		hash = gin.H{}
		for i := 0; i < len(achievements); i++ {
			achievement := achievements[i]
			hash[strconv.Itoa(achievement.Id)] = gin.H{
				"name":        achievement.Name,
				"description": achievement.Description,
				"reward":      achievement.Reward,
			}
		}
	} else {
		status = http.StatusBadRequest
		hash = nil
	}

	context.JSON(status, hash)
}
