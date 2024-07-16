package controllers

import (
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/database/models"
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
	id := context.Request.URL.Query().Get("id")
	var status int
	var user models.User
	if id == "" {
		var users []*models.User
		hash := gin.H{}
		userController.database.Find(&users)
		if len(users) > 0 {
			status = http.StatusOK
			for i := 0; i < len(users); i++ {
				user := users[i]
				hash[strconv.Itoa(user.Id)] = gin.H{
					"name":    user.Name,
					"surname": user.Surname,
					"balance": user.Balance,
				}
			}
		} else {
			status = http.StatusBadRequest
		}

		context.JSON(status, hash)
	} else {
		var users []models.User
		userController.database.
			Model(user).
			Where("id = ?", id).
			Find(&users)
		if len(users) == 0 {
			status = http.StatusBadRequest
		} else {
			status = http.StatusOK
			userController.database.Model(user).Where("id = ?", id).First(&user)
		}
		context.JSON(status, gin.H{
			"id":      user.Id,
			"name":    user.Name,
			"surname": user.Surname,
			"balance": user.Balance,
		})
	}
}

func (userController *UserController) Post(context *gin.Context) {
	query := context.Request.URL.Query()
	var user models.User
	var users []models.User
	var status int
	userController.database.
		Model(user).
		Where("username = ?", query.Get("username")).
		Find(&users)
	if len(users) == 0 {
		user = models.User{
			Username: query.Get("username"),
			Name:     query.Get("name"),
			Surname:  query.Get("surname"),
			Balance:  0,
		}
		userController.database.Save(&user)
		var role models.Role
		var roles []models.Role
		userController.database.
			Model(role).
			Where("key = ?", query.Get("role")).
			Find(&roles)
		if len(roles) == 1 {
			userController.database.Model(&role).Where("key = ?", query.Get("role")).First(&role)
		} else {
			userController.database.Model(&role).Where("key = ?", "user").First(&role)
		}
		role.Users = append(role.Users, user)
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(status, gin.H{
		"id":     user.Id,
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
