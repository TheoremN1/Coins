package migrations

import (
	"github.com/TheoremN1/Coins/database/models"
	"gorm.io/gorm"
)

func MigrationDown(database *gorm.DB) {
	migrator := database.Migrator()

	models := []interface{}{
		&models.CoinRequest{},
		&models.MerchRequest{},
		&models.User{},
		&models.Achievement{},
		&models.Product{},
		&models.Role{},
		&models.Status{},
	}

	for i := 0; i < len(models); i++ {
		model := models[i]
		if migrator.HasTable(model) {
			migrator.DropTable(model)
		}
	}
}
