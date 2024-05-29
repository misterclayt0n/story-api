package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system
// @Description User represents a user in the system
type User struct {
	gorm.Model
	// Username of the user
	Username string `json:"username" gorm:"unique"`
	// Password of the user
	Password string `json:"password"`
}
