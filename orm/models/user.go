package models

import "github.com/jinzhu/gorm"

// User describes users table
type User struct {
	gorm.Model
	Lastname  string
	Firstname string
	Username  string
}
