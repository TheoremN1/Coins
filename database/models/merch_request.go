package models

import "gorm.io/gorm"

type MerchRequest struct {
	gorm.Model
	Id          int `gorm:"primaryKey"`
	UserComment string
	HrComment   string
}
