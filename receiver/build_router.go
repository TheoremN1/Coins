package receiver

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/TheoremN1/Coins/receiver/configs"
	"github.com/TheoremN1/Coins/receiver/controllers"
	"github.com/TheoremN1/Coins/receiver/database/migrations"
	"github.com/TheoremN1/Coins/receiver/services"
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildRouter() *Router {
	// Configs
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

	// Database
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

	// Services
	rolesService := services.NewRolesService(database)
	usersService := services.NewUsersService(database, rolesService)

	// Controllers
	balanceController := controllers.NewBalanceController(database)
	userController := controllers.NewUserController(usersService)
	achievementsController := controllers.NewAchievementsController(database)
	requestController := controllers.NewRequestController(database, usersService)

	// Router
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
