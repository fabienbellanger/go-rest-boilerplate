package models

import "time"

// DefaultModel overrights gorm.Model definition
type DefaultModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TimestampModel type
type TimestampModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// IDModel type
type IDModel struct {
	ID uint `gorm:"primary_key"`
}
