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

func (userController *UserController) IsUsernameFree(username string) bool {
	var users []models.User
	userController.database.
		Where("username = ?", username).
		Find(&users)
	return len(users) == 0
}

func (userController *UserController) IsRoleExist(key string) bool {
	var roles []models.Role
	userController.database.
		Where("key = ?", key).
		Find(&roles)
	return len(roles) == 1
}

func (userController *UserController) IsUserExist(id string) bool {
	var users []models.User
	userController.database.
		Where("id = ?", id).
		Find(&users)
	return len(users) == 1
}

func (userController *UserController) Get(context *gin.Context) {
	id := context.Request.URL.Query().Get("id")
	var status int
	var hash gin.H
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
	} else {
		var user models.User
		if userController.IsUserExist(id) {
			status = http.StatusOK
			userController.database.
				Where("id = ?", id).
				First(&user)
			hash = gin.H{
				"id":      user.Id,
				"name":    user.Name,
				"surname": user.Surname,
				"balance": user.Balance,
			}
		} else {
			status = http.StatusBadRequest
			hash = gin.H{}
		}

	}
	context.JSON(status, hash)
}

func (userController *UserController) Post(context *gin.Context) {
	query := context.Request.URL.Query()
	var status int
	var hash gin.H
	if userController.IsUsernameFree(query.Get("username")) &&
		userController.IsRoleExist(query.Get("role")) {
		var user models.User
		var role models.Role
		user = models.User{
			Username: query.Get("username"),
			Name:     query.Get("name"),
			Surname:  query.Get("surname"),
			Balance:  0,
		}
		userController.database.Save(&user)
		userController.database.
			Where("key = ?", query.Get("role")).
			First(&role)
		role.Users = append(role.Users, user)
		status = http.StatusOK
		hash = gin.H{"id": user.Id}
	} else {
		status = http.StatusBadRequest
		hash = gin.H{}
	}
	context.JSON(status, hash)
}

func (userController *UserController) Put(context *gin.Context) {
	query := context.Request.URL.Query()
	var status int
	if userController.IsUserExist(query.Get("id")) {
		var user models.User
		userController.database.
			Where("id = ?", query.Get("id")).
			First(&user)
		user.Name = query.Get("name")
		user.Surname = query.Get("surname")
		userController.database.Save(user)
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(status, gin.H{})
}

// TODO: вроде как из контекста удаляет, но в БД пользователь остался
func (userController *UserController) Delete(context *gin.Context) {
	query := context.Request.URL.Query()
	var status int
	if userController.IsUserExist(query.Get("id")) {
		var user models.User
		userController.database.
			Where("id = ?", query.Get("id")).
			First(&user)
		userController.database.Delete(user)
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(status, gin.H{})
}
