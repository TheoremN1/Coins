package controllers

import "github.com/gin-gonic/gin"

type RequestController struct {
	databaseUrl string
}

func NewRequestController(databaseUrl string) *RequestController {
	return &RequestController{databaseUrl}
}

func (requestController *RequestController) Get(context *gin.Context) {

}

func (requestController *RequestController) Post(context *gin.Context) {

}

func (requestController *RequestController) Delete(context *gin.Context) {

}
