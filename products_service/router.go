package products_service

import (
	"encoding/json"
	"io"
	"os"

	"github.com/TheoremN1/Coins/ProductsService/configs"
	"github.com/TheoremN1/Coins/ProductsService/controllers"
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
	serverUrl := ":" + os.Getenv("PRODUCTS_SERVICE_PORT")
	databaseUrl := "database_service:" + os.Getenv("DATABASE_SERVICE_PORT")
	//reactUrl := GetUrl(filepath.Join("configs", "react.json"))

	engine := gin.Default()
	engine.Use(cors.Default())

	achievementsController := controllers.NewAchievementsController("http://" + databaseUrl)
	engine.GET("/api/achievements", achievementsController.Get)
	engine.POST("/api/achievements", achievementsController.Post)
	engine.PUT("/api/achievements", achievementsController.Put)
	engine.DELETE("/api/achievements", achievementsController.Delete)

	merchController := controllers.NewMerchController("http://" + databaseUrl)
	engine.GET("/api/merch", merchController.Get)
	engine.POST("/api/merch", merchController.Post)
	engine.PUT("/api/merch", merchController.Put)
	engine.DELETE("/api/merch", merchController.Delete)

	return &Router{engine, serverUrl}
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
