package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestController struct {
	database *gorm.DB
}

func NewRequestController(database *gorm.DB) *RequestController {
	return &RequestController{database}
}

func (requestController *RequestController) Get(context *gin.Context) {

}

func (requestController *RequestController) Post(context *gin.Context) {

}

func (requestController *RequestController) Delete(context *gin.Context) {

}
