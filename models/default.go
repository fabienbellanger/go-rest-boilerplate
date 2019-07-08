package models

import "time"

// PrimaryModel type
type PrimaryModel struct {
	ID uint64 `gorm:"primary_key" json:"id"`
}

// TimestampModel type
type TimestampModel struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// SoftDeleteModel type
type SoftDeleteModel struct {
	DeletedAt *time.Time `json:"deletedAt"` // * <=> nullable
}
