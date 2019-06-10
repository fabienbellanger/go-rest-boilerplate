package models

// Role describes roles table
type Role struct {
	PrimaryModel
	Label    string `gorm:"type:varchar(63);not null;"`
	ParentID uint   `gorm:"ForeignKey:Parent"`
	Parent   *Role  `gorm:"ForeignKey:ParentID"`
	Users    []User `gorm:"many2many:users_roles;"`
	TimestampModel
	SolfDeleteModel
}
