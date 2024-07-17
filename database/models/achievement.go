package models

import "gorm.io/gorm"

type Achievement struct {
	gorm.Model
	Id           int `gorm:"primaryKey"`
	Name         string
	Description  string
	Reward       int
	CoinRequests []CoinRequest `gorm:"foreignKey:Id"`
}
