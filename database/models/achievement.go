package models

import "gorm.io/gorm"

type Achievement struct {
	gorm.Model
	Id           int `gorm:"primaryKey"`
	Name         string
	Description  string
	Prize        int
	CoinRequests []CoinRequest `gorm:"foreignKey:Id"`
}
