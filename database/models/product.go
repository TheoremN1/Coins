package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id            int `gorm:"primaryKey"`
	Name          string
	Description   string
	Price         int
	FilePath      string
	MerchRequests []MerchRequest `gorm:"foreignKey:Id"`
}
