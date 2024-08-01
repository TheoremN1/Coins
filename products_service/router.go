package products_service

import (
	"os"

	"github.com/TheoremN1/Coins/ProductsService/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	url    string
}

func NewRouter() *Router {
	serverUrl := ":" + os.Getenv("PRODUCTS_SERVICE_PORT")
	databaseUrl := "http://nginx_coins/services/database"

	engine := gin.Default()
	engine.Use(cors.Default())

	achievementsController := controllers.NewAchievementsController(databaseUrl)
	engine.GET("/api/achievements", achievementsController.Get)
	engine.POST("/api/achievements", achievementsController.Post)
	engine.PUT("/api/achievements", achievementsController.Put)
	engine.DELETE("/api/achievements", achievementsController.Delete)

	merchController := controllers.NewMerchController(databaseUrl)
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
