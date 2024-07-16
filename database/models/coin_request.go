package models

import "gorm.io/gorm"

type CoinRequest struct {
	gorm.Model
	Id          int `gorm:"primaryKey"`
	UserComment string
	HrComment   string
	User        User        `gorm:"foreignKey:Id"`
	Status      Status      `gorm:"foreignKey:Id"`
	Achievement Achievement `gorm:"foreignKey:Id"`
}
