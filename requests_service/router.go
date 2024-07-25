package requests_service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/RequestsService/configs"
	"github.com/TheoremN1/Coins/RequestsService/controllers"
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

	engine := gin.Default()
	/*
		engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://" + serverUrl},
			AllowCredentials: true,
			/*
				AllowOriginFunc: func(origin string) bool {
					return origin == <URL на ReactApp>
				},
			/
			MaxAge: 12 * time.Hour,
		}))
	*/
	requestsController := controllers.NewRequestController("http://" + databaseUrl)
	engine.GET("/api/requests", requestsController.Get)
	engine.POST("/api/requests", requestsController.Post)
	engine.PUT("/api/requests", requestsController.Put)
	engine.DELETE("/api/requests", requestsController.Delete)

	return &Router{engine, serverUrl}
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
