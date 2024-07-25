package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/TheoremN1/Coins/UsersService/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	databaseUrl string
}

func NewRoleController(databaseUrl string) *RoleController {
	return &RoleController{databaseUrl}
}

func (roleController *RoleController) Get(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := roleController.databaseUrl + "/api/users/" + id + "/role"

	client := &http.Client{}
	req, err := sling.New().Get(url).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var role *models.Role
	json.NewDecoder(resp.Body).Decode(&role)

	context.JSON(resp.StatusCode, role)
}

func (roleController *RoleController) Put(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")
	role := query.Get("role")

	url := roleController.databaseUrl + "/api/users/" + id
	user, status := GetUser(url)

	if user == nil {
		context.JSON(status, false)
	} else {
		user.RoleKey = role
		isPuted, status := EditUser(url, user)
		context.JSON(status, isPuted)
	}
}
