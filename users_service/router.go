package users_service

import (
	"encoding/json"
	"io"
	"os"

	"github.com/TheoremN1/Coins/UsersService/configs"
	"github.com/TheoremN1/Coins/UsersService/controllers"
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
	serverUrl := ":" + os.Getenv("USERS_SERVICE_PORT")
	databaseUrl := "database_service:" + os.Getenv("DATABASE_SERVICE_PORT")
	//reactUrl := GetUrl(filepath.Join("configs", "react.json"))

	engine := gin.Default()
	engine.Use(cors.Default())

	userController := controllers.NewUserController("http://" + databaseUrl)
	engine.GET("/api/users", userController.Get)
	engine.POST("/api/users", userController.Post)
	engine.PUT("/api/users", userController.Put)
	engine.DELETE("/api/users", userController.Delete)

	balanceController := controllers.NewBalanceController("http://" + databaseUrl)
	engine.GET("/api/balance", balanceController.Get)
	engine.PUT("/api/balance", balanceController.Put)

	roleController := controllers.NewRoleController("http://" + databaseUrl)
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
