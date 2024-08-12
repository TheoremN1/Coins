package database

import (
	"fmt"

	"github.com/TheoremN1/Coins/auth/envs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	env := envs.GetDatabaseEnv()
	uri := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		env.Host, env.Port, env.Name, env.User, env.Password)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDatabase() *gorm.DB {
	return database
}
