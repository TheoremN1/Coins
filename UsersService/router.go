package users_service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/UsersService/configs"
	"github.com/TheoremN1/Coins/UsersService/controllers"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	url    string
}

func NewRouter() *Router {
	confFile, _ := os.Open(filepath.Join("configs", "server.json"))
	defer confFile.Close()
	bytes, _ := io.ReadAll(confFile)
	configuration := configs.ServerConfiguration{}
	json.Unmarshal(bytes, &configuration)

	balanceController := controllers.NewBalanceController()
	userController := controllers.NewUserController()

	router := Router{gin.Default(), configuration.Host + ":" + configuration.Port}

	router.engine.GET("/balance", balanceController.Get)
	router.engine.PUT("/balance", balanceController.Put)

	router.engine.GET("/users", userController.Get)
	router.engine.POST("/users", userController.Post)
	router.engine.PUT("/users", userController.Put)
	router.engine.DELETE("/users", userController.Delete)

	return &router
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
