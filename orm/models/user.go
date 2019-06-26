package models

// User describes users table
type User struct {
	PrimaryModel
	Username  string `gorm:"type:varchar(191);unique_index:idx_username" json:"username"`
	Password  string `gorm:"type:varchar(128);not null" json:"password"` // SHA-512
	Lastname  string `gorm:"type:varchar(100);not null" json:"lastname"`
	Firstname string `gorm:"type:varchar(100);not null" json:"firstname"`
	// Roles     []Role `gorm:"many2many:users_roles;" json:"roles"`
	TimestampModel
	SoftDeleteModel
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
