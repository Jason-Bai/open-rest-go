package models

import "time"

// Base base model definition
type Base struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
