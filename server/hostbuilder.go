package server

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/configs"
	"github.com/TheoremN1/Coins/server/controllers"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Host struct {
	Database          *gorm.DB
	BalanceController *controllers.BalanceController
	Router            *Router
}

func NewHost() *Host {
	confFile, err := os.Open(filepath.Join("configs", "config.json"))
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

	host := Host{}
	stringConnection := "host=" + conf.Database.Host +
		" user=" + conf.Database.Username +
		" dbname=" + conf.Database.Name +
		" password=" + conf.Database.Password
	database, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	host.Database = database
	host.BalanceController = controllers.NewBalanceController()

	host.Router = NewRouter(
		gin.Default(),
		host.BalanceController,
		conf.Server.Host+":"+conf.Server.Port,
	)

	return &host
}
