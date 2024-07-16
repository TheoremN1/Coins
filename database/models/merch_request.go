package models

import "gorm.io/gorm"

type MerchRequest struct {
	gorm.Model
	Id          int `gorm:"primaryKey"`
	UserComment string
	HrComment   string
	User        User   `gorm:"foreignKey:Id"`
	Status      Status `gorm:"foreignKey:Id"`
	Product     Product  `gorm:"foreignKey:Id"`
}
