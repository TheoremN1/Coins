package internal

import (
	"github.com/TheoremN1/Coins/internal/controllers"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	Run()
}

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) IRouter {
	return &Router{engine}
}

func (router *Router) Run() {
	healthController := controllers.NewHealthController()
	indexController := controllers.NewIndexController()

	router.engine.GET("/", indexController.Index)
	router.engine.GET("/check", healthController.Check)

	// listen 0.0.0.0:8080 (localhost:8080)
	router.engine.Run()
}
