package role

import (
	defaultModel "github.com/fabienbellanger/go-rest-boilerplate/models/orm"
	userModel "github.com/fabienbellanger/go-rest-boilerplate/models/user"
)

// Role describes roles table
type Role struct {
	defaultModel.PrimaryModel
	Label    string `gorm:"type:varchar(63);not null;" json:"label"`
	ParentID uint   `json:"parentId"`
	Parent   *Role
	Users    []userModel.User `gorm:"many2many:users_roles;" json:"users"`
	defaultModel.TimestampModel
	defaultModel.SoftDeleteModel
}
