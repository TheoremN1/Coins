package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Id            int    `gorm:"primaryKey"`
	Key           string `gorm:"unique"`
	Name          string
	CoinRequests  []CoinRequest  `gorm:"foreignKey:Id"`
	MerchRequests []MerchRequest `gorm:"foreignKey:Id"`
}
