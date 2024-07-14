package controllers

import "github.com/gin-gonic/gin"

type IHealthController interface {
	Check(context *gin.Context)
}

type HealthController struct{}

func (healthController *HealthController) Check(context *gin.Context) {
	context.JSON(200, gin.H{"status": "ok"})
}

func NewHealthController() IHealthController {
	return &HealthController{}
}
