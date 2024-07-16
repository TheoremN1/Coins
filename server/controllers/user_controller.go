package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	database *gorm.DB
}

func NewUserController(database *gorm.DB) *UserController {
	return &UserController{database}
}

func (userController *UserController) Get(context *gin.Context) {
	status := http.StatusOK
	context.JSON(status, gin.H{
		"method": "get",
		"status": status,
	})
}

func (userController *UserController) Post(context *gin.Context) {
	status := http.StatusOK
	context.JSON(status, gin.H{
		"method": "post",
		"status": status,
	})
}

func (userController *UserController) Put(context *gin.Context) {
	status := http.StatusOK
	context.JSON(status, gin.H{
		"method": "put",
		"status": status,
	})
}

func (userController *UserController) Delete(context *gin.Context) {
	status := http.StatusOK
	context.JSON(status, gin.H{
		"method": "delete",
		"status": status,
	})
}
