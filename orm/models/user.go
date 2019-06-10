package models

// User describes users table
type User struct {
	PrimaryModel
	Username  string `gorm:"type:varchar(191);unique_index:index_users_username"`
	Password  string `gorm:"type:varchar(128;not null"` // SHA-512
	Lastname  string `gorm:"type:varchar(100);not null"`
	Firstname string `gorm:"type:varchar(100);not null"`
	Roles     []Role `gorm:"many2many:users_roles;"`
	TimestampModel
	SolfDeleteModel
}

// GetFullname returns user fullname
func (u *User) GetFullname() string {
	if u.Firstname == "" {
		return u.Lastname
	} else if u.Lastname == "" {
		return u.Firstname
	}

	return u.Firstname + " " + u.Lastname
}
