package server

import (
	"encoding/json"
	"io"
	"os"

	"github.com/TheoremN1/Coins/configs"
	"github.com/TheoremN1/Coins/server/controllers"
	"github.com/gin-gonic/gin"
)

type Host struct {
	Router *Router
}

func NewHost() *Host {
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
	indexController := controllers.NewIndexController()
	healthController := controllers.NewHealthController()

	router := NewRouter(
		gin.Default(),
		indexController,
		healthController,
		conf.Server.Host+":"+conf.Server.Port,
	)

	return &Host{router}
}
