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
	// Input
	query := context.Request.URL.Query()
	id := query.Get("id")
	// Output
	var status int
	var hash gin.H

	if id == "" {
		// All users
		var users []*models.User
		hash = gin.H{}
		userController.database.Model(&models.User{}).Find(&users)
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
			hash = nil
		}
	} else {
		// Only one user
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
			hash = nil
		}
	}

	context.JSON(status, hash)
}

func (userController *UserController) Post(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	username := query.Get("username")
	name := query.Get("name")
	surname := query.Get("surname")
	role := query.Get("role")
	// Output
	var status int
	var hash gin.H

	if userController.IsUsernameFree(username) &&
		userController.IsRoleExist(role) {
		// Username is free and role is exist
		user := models.User{
			Username: username,
			Name:     name,
			Surname:  surname,
			Balance:  0,
		}
		userController.database.Save(&user)

		var role models.Role
		userController.database.
			Where("key = ?", role).
			First(&role)
		role.Users = append(role.Users, user) // TODO: Не уверен что в БД записывается связь между юзером и ролью

		status = http.StatusOK
		hash = gin.H{"id": user.Id}
	} else {
		status = http.StatusBadRequest
		hash = nil
	}

	context.JSON(status, hash)
}

func (userController *UserController) Put(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	id := query.Get("id")
	name := query.Get("name")
	surname := query.Get("surname")
	// Output
	var status int

	if userController.IsUserExist(id) {
		var user models.User
		userController.database.
			Where("id = ?", id).
			First(&user)
		user.Name = name
		user.Surname = surname
		userController.database.Save(user)
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(status, nil)
}

func (userController *UserController) Delete(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	id := query.Get("id")
	// Output
	var status int

	if userController.IsUserExist(id) {
		var user models.User
		userController.database.
			Where("id = ?", id).
			First(&user)
		userController.database.Delete(&user)
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}

	context.JSON(status, nil)
}
