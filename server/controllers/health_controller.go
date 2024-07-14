package controllers

import "github.com/gin-gonic/gin"

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (healthController *HealthController) Get(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}
