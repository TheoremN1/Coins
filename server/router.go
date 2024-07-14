package server

import (
	"encoding/json"
	"io"
	"os"

	"github.com/TheoremN1/Coins/configs"
	"github.com/TheoremN1/Coins/server/controllers"

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

	confFile, err := os.Open("configs/config.json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	bytes, err := io.ReadAll(confFile)
	if err != nil {
		panic(err)
	}
	conf := configs.LaunchConf{}
	err = json.Unmarshal(bytes, &conf)
	if err != nil {
		panic(err)
	}
	err = router.engine.Run(conf.Server.Host + ":" + conf.Server.Port)
	if err != nil {
		panic(err)
	}
}
