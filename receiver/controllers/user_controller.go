package controllers

import (
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/receiver/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	usersService *services.UsersService
}

func NewUserController(usersService *services.UsersService) *UserController {
	return &UserController{usersService}
}

func (userController *UserController) Get(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	// with "id" - if need only one user
	// without "id" - if need all users

	// Output
	var status int
	var hash gin.H

	if query.Has("id") {
		// Only one user
		idStr := query.Get("id")
		id, err := strconv.Atoi(idStr)
		if err == nil {
			user := userController.usersService.GetUser(id)
			if user != nil {
				hash = gin.H{
					"id":      user.Id,
					"name":    user.Name,
					"surname": user.Surname,
					"balance": user.Balance,
				}
				status = http.StatusOK
			} else {
				status = http.StatusBadRequest
				hash = nil
			}
		} else {
			status = http.StatusBadRequest
			hash = nil
		}
	} else {
		// All users
		users := userController.usersService.GetAllUsers()
		if len(users) > 0 {
			status = http.StatusOK
			hash = gin.H{}
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

	isCreated, id := userController.usersService.NewUser(username, name, surname, role)

	if isCreated {
		status = http.StatusOK
		hash = gin.H{"id": id}
	} else {
		status = http.StatusBadRequest
		hash = nil
	}

	context.JSON(status, hash)
}

func (userController *UserController) Put(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	idStr := query.Get("id")
	name := query.Get("name")
	surname := query.Get("surname")
	// Output
	var status int

	id, err := strconv.Atoi(idStr)
	if err == nil {
		isEdited := userController.usersService.EditUser(id, name, surname)
		if isEdited {
			status = http.StatusOK
		} else {
			status = http.StatusBadRequest
		}
	} else {
		status = http.StatusBadRequest
	}

	context.JSON(status, nil)
}

func (userController *UserController) Delete(context *gin.Context) {
	// Input
	query := context.Request.URL.Query()
	idStr := query.Get("id")
	// Output
	var status int
	id, err := strconv.Atoi(idStr)
	if err == nil {
		isDeleted := userController.usersService.DeleteUser(id)
		if isDeleted {
			status = http.StatusOK
		} else {
			status = http.StatusBadRequest
		}
	} else {
		status = http.StatusBadRequest
	}

	context.JSON(status, nil)
}
