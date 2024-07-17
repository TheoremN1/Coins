package server

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/configs"
	"github.com/TheoremN1/Coins/database/migrations"
	"github.com/TheoremN1/Coins/server/controllers"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildRouter() *Router {
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

	stringConnection := "host=" + conf.Database.Host +
		" user=" + conf.Database.Username +
		" dbname=" + conf.Database.Name +
		" password=" + conf.Database.Password
	database, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migrations.MigrationDown(database)
	migrations.MigrationUp(database)

	balanceController := controllers.NewBalanceController(database)
	userController := controllers.NewUserController(database)
	achievementsController := controllers.NewAchievementsController(database)
	requestController := controllers.NewRequestController(database)
	router := NewRouter(
		gin.Default(),
		balanceController,
		userController,
		requestController,
		achievementsController,
		conf.Server.Host+":"+conf.Server.Port,
	)

	return router
}
