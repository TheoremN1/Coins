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
	balanceController *controllers.BalanceController,
	url string,
) *Router {
	router := Router{engine, url}
	router.engine.GET("/balance", balanceController.Get)
	return &router
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
