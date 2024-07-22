package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (userController *UserController) Get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"type": "get"})
}

func (userController *UserController) Post(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"type": "post"})
}

func (userController *UserController) Put(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"type": "put"})
}

func (userController *UserController) Delete(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"type": "delete"})
}
