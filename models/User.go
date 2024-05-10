package models

import "gorm.io/gorm"

type User struct {
	// esto lee nuestro struct y los convierte en una tabla
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	Email     string `gorm:"not null:unique_index" json:"email"`
	Tasks     []Task `json:"task"`
}
