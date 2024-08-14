package controllers

import (
	"net/http"

	"github.com/TheoremN1/Coins/auth/database"
	"github.com/TheoremN1/Coins/auth/database/entities"
	"github.com/TheoremN1/Coins/auth/server/models"
	"github.com/TheoremN1/Coins/auth/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthController struct {
	database *gorm.DB
}

var authController *AuthController

func init() {
	database := database.GetDatabase()
	authController = &AuthController{database}
}

func GetAuthController() *AuthController {
	return authController
}

// Methods

func (ac *AuthController) Registration(ctx *gin.Context) {
	var registrationData models.RegistrationData
	if err := ctx.ShouldBindJSON(&registrationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := utils.HashPassword(registrationData.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user entities.User
	user.Id = uuid.New().String()
	user.Name = registrationData.Name
	user.Surname = registrationData.Surname
	user.Email = registrationData.Email
	user.Password = hash

	err = ac.database.Create(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (ac *AuthController) Authorization(ctx *gin.Context) {
	var authorizationData models.AuthorizationData
	if err := ctx.ShouldBindJSON(&authorizationData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user entities.User
	result := ac.database.Where("email = ?", authorizationData.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "неверный email или пароль"})
		return
	}
	if !utils.CheckPasswordHash(authorizationData.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "неверный email или пароль"})
		return
	}

	tokens, err := utils.GenerateTokens(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

func (*AuthController) RefreshToken(ctx *gin.Context) {
	var token models.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.ValidateRefreshToken(token.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokens, err := utils.GenerateTokens(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (_ *AuthController) ExtractData(ctx *gin.Context) {
	var token models.Tokens
	if err := ctx.ShouldBindJSON(&token); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := utils.ExtractDataFromAccessToken(token.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
