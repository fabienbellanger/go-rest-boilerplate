package models

// Role describes roles table
type Role struct {
	PrimaryModel
	Label    string `gorm:"type:varchar(63);not null;"`
	ParentID uint
	Parent   *Role
	// Users    []User `gorm:"many2many:users_roles;"`
	TimestampModel
	SoftDeleteModel
}
