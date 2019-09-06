package models

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// User describes users table
type User struct {
	PrimaryModel
	Username  string `gorm:"type:varchar(191);unique_index:idx_username" json:"username"`
	Password  string `gorm:"type:varchar(128);not null" json:"password"` // SHA-256
	Lastname  string `gorm:"type:varchar(100);not null" json:"lastname"`
	Firstname string `gorm:"type:varchar(100);not null" json:"firstname"`
	// Roles     []Role `gorm:"many2many:users_roles;" json:"roles"`
	TimestampModel
	SoftDeleteModel
}

// CheckLogin checks if username and password are corrects
func (u *User) CheckLogin(username, password string) error {
	if len(username) == 0 || len(password) == 0 {
		return nil
	}

	encryptPassword := sha512.Sum512([]byte(password))
	encryptPasswordStr := hex.EncodeToString(encryptPassword[:])
	query := `
		SELECT id, username, lastname, firstname, created_at, deleted_at
		FROM users
		WHERE username = ? AND password = ? AND deleted_at IS NULL
		LIMIT 1`
	rows, err := database.Select(query, username, encryptPasswordStr)

	for rows.Next() {
		println()
		err = rows.Scan(
			&u.ID,
			&u.Username,
			&u.Lastname,
			&u.Firstname,
			&u.CreatedAt,
			&u.DeletedAt)

		lib.CheckError(err, 0)
	}

	return err
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

// ChangePassword changes user password in database
func (u *User) ChangePassword(password string) bool {
	return len(database.Orm.Model(&u).Update("password", password).GetErrors()) == 0
}
