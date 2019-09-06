package module

import (
	applicationModel "github.com/fabienbellanger/go-rest-boilerplate/models/application"
	defaultModel "github.com/fabienbellanger/go-rest-boilerplate/models/orm"
)

// Module describes modules table
type Module struct {
	defaultModel.PrimaryModel
	Name          string `gorm:"type:varchar(191)" json:"name"`
	ApplicationID int
	Application   applicationModel.Application
	defaultModel.TimestampModel
}
