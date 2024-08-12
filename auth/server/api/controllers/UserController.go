package controllers

import (
	"net/http"

	"github.com/TheoremN1/Coins/auth/database"
	"github.com/TheoremN1/Coins/auth/database/entities"
	"github.com/TheoremN1/Coins/auth/server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (uc *UserController) GetId(ctx *gin.Context) {
	id := ctx.Param("id")
	var user entities.User
	err := uc.database.Where("Id = ?", id).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	var users []entities.User
	err := uc.database.Find(&users).Error
	if err != nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) Post(ctx *gin.Context) {
	var registrationData models.RegistrationData
	if err := ctx.ShouldBindJSON(&registrationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entities.User
	user.Id = uuid.New().String()
	user.Name = registrationData.Name
	user.Surname = registrationData.Surname
	user.Email = registrationData.Email
	// TODO: пароль по-хорошему надо шифровать перед записью в БД
	user.Password = registrationData.Password

	err := uc.database.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (uc *UserController) Put(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"method": "put"})
}

func (uc *UserController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, gin.H{"method": "delete"})
}
