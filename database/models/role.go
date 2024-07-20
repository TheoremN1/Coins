package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Id    int    `gorm:"primaryKey"`
	Key   string `gorm:"unique"`
	Name  string
	Users []User `gorm:"foreignKey:Id"`
}
