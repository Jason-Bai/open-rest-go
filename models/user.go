package models

import (
	"time"
)

// User a model of user
type User struct {
	Base
	Name     string
	Age      uint
	Birthday *time.Time
	Email    string `gorm:"type:varchar(100);unique_index"`
	Role     string `gorm:"size:255"` // set field size to 255        // ignore this field
}

// TableName rename user's table name
func (User) TableName() string {
	return "user"
}
