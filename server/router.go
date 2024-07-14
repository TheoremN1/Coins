package server

import (
	"github.com/TheoremN1/Coins/server/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	url    string
}

func NewRouter(
	engine *gin.Engine,
	indexController *controllers.IndexController,
	healthController *controllers.HealthController,
	url string,
) *Router {
	router := Router{engine, url}
	router.engine.GET("/", indexController.Get)
	router.engine.GET("/check", healthController.Get)
	return &router
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
