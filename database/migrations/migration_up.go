package migrations

import (
	"github.com/TheoremN1/Coins/database/models"
	"gorm.io/gorm"
)

func MigrationUp(database *gorm.DB) {
	database.AutoMigrate(
		&models.Achievement{},
		&models.Product{},
		&models.Role{},
		&models.Status{},
		&models.User{},
		&models.CoinRequest{},
		&models.MerchRequest{},
	)

	roles := []*models.Role{
		{Key: "user", Name: "Пользователь"},
		{Key: "hr", Name: "HR"},
		{Key: "admin", Name: "Администратор"},
	}
	database.Save(&roles)

	statuses := []*models.Status{
		{Key: "wait", Name: "Ожидание"},
		{Key: "ready", Name: "Готово"},
		{Key: "denied", Name: "Отказано"},
	}
	database.Save(&statuses)
}
