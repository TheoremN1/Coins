package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id            int    `gorm:"primaryKey"`
	Username      string `gorm:"unique"`
	Name          string
	Surname       string
	Balance       int
	CoinRequests  []CoinRequest  `gorm:"foreignKey:Id"`
	MerchRequests []MerchRequest `gorm:"foreignKey:Id"`
}
