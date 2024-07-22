package users_service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/TheoremN1/Coins/UsersService/configs"
	"github.com/TheoremN1/Coins/UsersService/controllers"
	"github.com/gin-contrib/cors"
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
	url := configuration.Host + ":" + configuration.Port

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://" + url},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		/*
			AllowOriginFunc: func(origin string) bool {
				return origin == <URL на ReactApp>
			},
		*/
		MaxAge: 12 * time.Hour,
	}))

	userController := controllers.NewUserController()
	engine.GET("/users", userController.Get)
	engine.POST("/users", userController.Post)
	engine.PUT("/users", userController.Put)
	engine.DELETE("/users", userController.Delete)

	balanceController := controllers.NewBalanceController()
	engine.GET("/balance", balanceController.Get)
	engine.PUT("/balance", balanceController.Put)

	return &Router{engine, url}
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
