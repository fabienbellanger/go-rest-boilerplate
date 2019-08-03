package models

// Application describes applications table
type Application struct {
	PrimaryModel
	Name string `gorm:"type:varchar(191)" json:"name"`
	TimestampModel
}
