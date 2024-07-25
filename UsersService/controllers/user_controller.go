package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/TheoremN1/Coins/UsersService/models"
	"github.com/dghubble/sling"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	databaseUrl string
}

func NewUserController(databaseUrl string) *UserController {
	return &UserController{databaseUrl}
}

func GetUser(url string) (*models.User, int) {
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

	var user *models.User
	json.NewDecoder(resp.Body).Decode(&user)
	return user, resp.StatusCode
}

func (userController *UserController) Get(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := userController.databaseUrl + "/api/users/" + id

	user, status := GetUser(url)

	context.JSON(status, user)
}

func (userController *UserController) Post(context *gin.Context) {
	var user *models.User
	json.NewDecoder(context.Request.Body).Decode(&user)

	url := userController.databaseUrl + "/api/users"
	client := &http.Client{}
	req, err := sling.New().Post(url).BodyForm(user).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var isPosted bool
	json.NewDecoder(resp.Body).Decode(&isPosted)

	context.JSON(resp.StatusCode, isPosted)
}

func EditUser(url string, user *models.User) (bool, int) {
	client := &http.Client{}
	req, err := sling.New().Put(url).BodyForm(user).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var isPuted bool
	json.NewDecoder(resp.Body).Decode(&isPuted)

	return isPuted, resp.StatusCode
}

func (userController *UserController) Put(context *gin.Context) {
	var user *models.User
	json.NewDecoder(context.Request.Body).Decode(&user)

	url := userController.databaseUrl + "/api/users/" + strconv.Itoa(user.Id)

	isPuted, status := EditUser(url, user)

	context.JSON(status, isPuted)
}

func (userController *UserController) Delete(context *gin.Context) {
	query := context.Request.URL.Query()
	id := query.Get("id")

	url := userController.databaseUrl + "/api/users/" + id
	client := &http.Client{}
	req, err := sling.New().Delete(url).Request()
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var isDeleted *bool
	json.NewDecoder(resp.Body).Decode(&isDeleted)

	context.JSON(resp.StatusCode, isDeleted)
}
