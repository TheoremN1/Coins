package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	databaseUrl string
}

func NewUserController(databaseUrl string) *UserController {
	return &UserController{databaseUrl}
}

func (userController *UserController) Get(context *gin.Context) {
	responce, _ := http.Get(userController.databaseUrl + "/api/users")
	var j interface{}
	json.NewDecoder(responce.Body).Decode(&j)
	context.JSON(http.StatusOK, j)
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
