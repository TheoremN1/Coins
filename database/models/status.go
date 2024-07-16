package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string
}
