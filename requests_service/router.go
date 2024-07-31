package requests_service

import (
	"os"

	"github.com/TheoremN1/Coins/RequestsService/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	url    string
}

func NewRouter() *Router {
	serverUrl := ":" + os.Getenv("REQUEST_SERVICE_PORT")
	databaseUrl := "http://database_service"

	engine := gin.Default()
	engine.Use(cors.Default())

	coinsRequestsController := controllers.NewCoinsRequestController(databaseUrl)
	engine.GET("/api/coinsrequests", coinsRequestsController.Get)
	engine.POST("/api/coinsrequests", coinsRequestsController.Post)
	engine.PUT("/api/coinsrequests", coinsRequestsController.Put)
	engine.DELETE("/api/coinsrequests", coinsRequestsController.Delete)

	merchRequestsController := controllers.NewMerchRequestController(databaseUrl)
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
