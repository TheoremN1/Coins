package users_service

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/UsersService/configs"
	"github.com/TheoremN1/Coins/UsersService/controllers"
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
	serverUrl := GetUrl(filepath.Join("configs", "server.json"))
	databaseUrl := GetUrl(filepath.Join("configs", "database.json"))

	engine := gin.Default()
	/*
		engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://" + serverUrl},
			AllowCredentials: true,
			/*
				AllowOriginFunc: func(origin string) bool {
					return origin == <URL на ReactApp>
				},
			/
			MaxAge: 12 * time.Hour,
		}))
	*/
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
