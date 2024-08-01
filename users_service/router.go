package users_service

import (
	"os"

	"github.com/TheoremN1/Coins/UsersService/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
	url    string
}

func NewRouter() *Router {
	serverUrl := ":" + os.Getenv("USERS_SERVICE_PORT")
	databaseUrl := "http://nginx_coins/services/database"

	engine := gin.Default()
	engine.Use(cors.Default())

	userController := controllers.NewUserController(databaseUrl)
	engine.GET("/api/users", userController.Get)
	engine.POST("/api/users", userController.Post)
	engine.PUT("/api/users", userController.Put)
	engine.DELETE("/api/users", userController.Delete)

	balanceController := controllers.NewBalanceController(databaseUrl)
	engine.GET("/api/balance", balanceController.Get)
	engine.PUT("/api/balance", balanceController.Put)

	roleController := controllers.NewRoleController(databaseUrl)
	engine.GET("/api/role", roleController.Get)
	engine.PUT("/api/role", roleController.Put)

	return &Router{engine, serverUrl}
}

func (router *Router) Run() {
	err := router.engine.Run(router.url)
	if err != nil {
		panic(err)
	}
}
