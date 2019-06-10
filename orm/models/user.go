package models

// User describes users table
type User struct {
	Model
	Username  string `gorm:"type:varchar(191);unique_index:index_user_username"`
	Password  string `gorm:"type:varchar(128;not null"`
	Lastname  string `gorm:"type:varchar(100);not null"`
	Firstname string `gorm:"type:varchar(100);not null"`
}

// GetFullname returns user fullname
func (u *User) GetFullname() string {
	return u.Firstname + " " + u.Lastname
}
