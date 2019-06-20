package models

// Role describes roles table
type Role struct {
	PrimaryModel
	Label    string `gorm:"type:varchar(63);not null;" json:"label"`
	ParentID uint   `json:"parentId"`
	Parent   *Role
	Users    []User `gorm:"many2many:users_roles;" json:"users"`
	TimestampModel
	SoftDeleteModel
}
