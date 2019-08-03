package models

// Module describes modules table
type Module struct {
	PrimaryModel
	Name          string `gorm:"type:varchar(191)" json:"name"`
	ApplicationID int
	Application   Application
	TimestampModel
}
