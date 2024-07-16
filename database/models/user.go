package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id      int `gorm:"primaryKey"`
	Name    string
	Surname string
	Balance int
	Role    Role `gorm:"foreignKey:Id"`
}
