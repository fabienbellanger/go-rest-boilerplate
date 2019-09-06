package application

import defaultModel "github.com/fabienbellanger/go-rest-boilerplate/models/orm"

// Application describes applications table
type Application struct {
	defaultModel.PrimaryModel
	Name string `gorm:"type:varchar(191)" json:"name"`
	defaultModel.TimestampModel
}
