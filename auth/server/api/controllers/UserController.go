package controllers

import (
	"net/http"

	"github.com/TheoremN1/Coins/auth/database"
	"github.com/TheoremN1/Coins/auth/database/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	database *gorm.DB
}

var userController *UserController

func init() {
	database := database.GetDatabase()
	userController = &UserController{database}
}

func GetUserController() *UserController {
	return userController
}

// Methods

func (uc *UserController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	var user entities.User
	err := uc.database.Where("Id = ?", id).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	var users []entities.User
	err := uc.database.Find(&users).Error
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) DeleteUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	var user entities.User
	err := uc.database.Where("Id = ?", id).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	err = uc.database.Delete(user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
