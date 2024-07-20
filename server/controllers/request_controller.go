package controllers

import (
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/server/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestController struct {
	database     *gorm.DB
	usersService *services.UsersService
}

func NewRequestController(database *gorm.DB, usersService *services.UsersService) *RequestController {
	return &RequestController{database, usersService}
}

func (requestController *RequestController) Get(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	idStr := query.Get("id")
	typeOfRequest := query.Get("type")

	// Output
	var status int
	var hash gin.H

	id, err := strconv.Atoi(idStr)
	if err == nil {
		user := requestController.usersService.GetUser(id)
		if user != nil {
			switch typeOfRequest {
			case "coins":
				status = http.StatusOK
				hash = gin.H{}
				// TODO: Выводить инфу про достижение
				for _, request := range user.CoinRequests {
					hash[strconv.Itoa(request.Id)] = gin.H{
						"userCom": request.UserComment,
						"hrCom":   request.HrComment,
					}
				}
			case "merch":
				status = http.StatusOK
				hash = gin.H{}
				// TODO: Выводить инфу про мерч
				for _, request := range user.MerchRequests {
					hash[strconv.Itoa(request.Id)] = gin.H{
						"userCom": request.UserComment,
						"hrCom":   request.HrComment,
					}
				}
			default:
				status = http.StatusBadRequest
				hash = nil
			}
		} else {
			status = http.StatusBadRequest
			hash = nil
		}
	} else {
		status = http.StatusBadRequest
		hash = nil
	}

	context.JSON(status, hash)
}

func (requestController *RequestController) Post(context *gin.Context) {

}

func (requestController *RequestController) Delete(context *gin.Context) {

}
