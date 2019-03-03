package user

import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// UserDB type
type UserDB struct {
	ID        uint64
	Username  string
	Password  string
	Lastname  string
	Firstname string
	CreatedAt sql.RawBytes
	DeletedAt sql.RawBytes
}

// CheckLogin : Authentification
func CheckLogin(username, password string) (UserDB, error) {
	encryptPassword := sha512.Sum512([]byte(password))
	encryptPasswordStr := hex.EncodeToString(encryptPassword[:])
	query := `
		SELECT id, username, lastname, firstname, created_at, deleted_at
		FROM user
		WHERE username = ? AND password = ? AND deleted_at IS NULL
		LIMIT 1`
	rows, err := database.Select(query, username, encryptPasswordStr)

	var user UserDB

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Lastname, &user.Firstname, &user.CreatedAt, &user.DeletedAt)

		lib.CheckError(err, 0)
	}

	return user, err
}

// GetFullname : Affichage le nom complet
func (u UserDB) GetFullname() string {
	return u.Firstname + " " + u.Lastname
}
