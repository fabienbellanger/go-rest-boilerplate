package models

import "time"

// PrimaryModel type
type PrimaryModel struct {
	ID uint `gorm:"primary_key" json:"id"`
}

// TimestampModel type
type TimestampModel struct {
	CreatedAt time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null" json:"updatedAt"`
}

// SoftDeleteModel type
type SoftDeleteModel struct {
	DeletedAt *time.Time `json:"deletedAt"`
}
