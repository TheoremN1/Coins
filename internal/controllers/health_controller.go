package controllers

import "github.com/gin-gonic/gin"

type HealthController struct{}

func (healthController *HealthController) Check(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}

func NewHealthController() *HealthController {
	return &HealthController{}
}
