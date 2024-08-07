package products_service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/ProductsService/configs"
	"github.com/TheoremN1/Coins/ProductsService/controllers"
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
	achievementsController := controllers.NewAchievementsController("http://" + databaseUrl)
	engine.GET("/api/achievements", achievementsController.Get)
	engine.POST("/api/achievements", achievementsController.Post)
	engine.PUT("/api/achievements", achievementsController.Put)
	engine.DELETE("/api/achievements", achievementsController.Delete)

	return &Router{engine, serverUrl}
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
