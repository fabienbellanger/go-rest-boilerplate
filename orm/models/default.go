package models

import "time"

// PrimaryModel type
type PrimaryModel struct {
	ID uint `gorm:"primary_key"`
}

// TimestampModel type
type TimestampModel struct {
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// SolfDeleteModel type
type SolfDeleteModel struct {
	DeletedAt *time.Time
}
