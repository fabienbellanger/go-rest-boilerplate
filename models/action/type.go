package action

import (
	moduleModel "github.com/fabienbellanger/go-rest-boilerplate/models/module"
	defaultModel "github.com/fabienbellanger/go-rest-boilerplate/models/orm"
)

// Action describes actions table
type Action struct {
	defaultModel.PrimaryModel
	Name     string `gorm:"type:varchar(191)" json:"name"`
	ModuleID int
	Module   moduleModel.Module
	defaultModel.TimestampModel
}
