package models

// Action describes actions table
type Action struct {
	PrimaryModel
	Name     string `gorm:"type:varchar(191)" json:"name"`
	ModuleID int
	Module   Module
	TimestampModel
}
