package requests_service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/RequestsService/configs"
	"github.com/TheoremN1/Coins/RequestsService/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	url    string
}

func GetUrl(jsonPath string) string {
	confFile, _ := os.Open(jsonPath)
	defer confFile.Close()
	bytes, _ := io.ReadAll(confFile)
	configuration := configs.ConnectionConfiguration{}
	json.Unmarshal(bytes, &configuration)
	return configuration.Host + ":" + configuration.Port
}

func NewRouter() *Router {
	serverUrl := GetUrl(filepath.Join("configs", "server.json"))
	databaseUrl := GetUrl(filepath.Join("configs", "database.json"))
	reactUrl := GetUrl(filepath.Join("configs", "react.json"))

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://" + reactUrl},
	}))

	coinsRequestsController := controllers.NewCoinsRequestController("http://" + databaseUrl)
	engine.GET("/api/coinsrequests", coinsRequestsController.Get)
	engine.POST("/api/coinsrequests", coinsRequestsController.Post)
	engine.PUT("/api/coinsrequests", coinsRequestsController.Put)
	engine.DELETE("/api/coinsrequests", coinsRequestsController.Delete)

	merchRequestsController := controllers.NewMerchRequestController("http://" + databaseUrl)
	engine.GET("/api/merchrequests", merchRequestsController.Get)
	engine.POST("/api/merchrequests", merchRequestsController.Post)
	engine.PUT("/api/merchrequests", merchRequestsController.Put)
	engine.DELETE("/api/merchrequests", merchRequestsController.Delete)

	return &Router{engine, serverUrl}
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
