package models

import "gorm.io/gorm"

type User struct {
	// esto lee nuestro struct y los convierte en una tabla
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null:unique_index"`
	Tasks     []Task `json:"task"`
}
