package receiver

import (
	"github.com/TheoremN1/Coins/receiver/controllers"

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
	requestController *controllers.RequestController,
	achievementsController *controllers.AchievementsController,
	url string,
) *Router {

	router := Router{engine, url}

	router.engine.GET("/balance", balanceController.Get)
	router.engine.PUT("/balance", balanceController.Put)

	router.engine.GET("/users", userController.Get)
	router.engine.POST("/users", userController.Post)
	router.engine.PUT("/users", userController.Put)
	router.engine.DELETE("/users", userController.Delete)

	router.engine.GET("/requests", requestController.Get)
	router.engine.POST("/requests", requestController.Post)
	router.engine.DELETE("/requests", requestController.Delete)

	router.engine.GET("/achievements", achievementsController.Get)

	return &router
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}