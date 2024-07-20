package models

import "gorm.io/gorm"

type CoinRequest struct {
	gorm.Model
	Id          int `gorm:"primaryKey"`
	UserComment string
	HrComment   string
}
