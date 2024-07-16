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
	userController *controllers.UserController,
	url string,
) *Router {
	router := Router{engine, url}
	router.engine.GET("/balance", balanceController.Get)
	router.engine.GET("/users", userController.Get)
	router.engine.POST("/users", userController.Post)
	router.engine.PUT("/users", userController.Put)
	router.engine.DELETE("/users", userController.Delete)
	return &router
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
