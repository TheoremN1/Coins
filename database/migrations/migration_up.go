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
}
