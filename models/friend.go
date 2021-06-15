package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}
