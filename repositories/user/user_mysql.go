package user

import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/fabienbellanger/go-rest-boilerplate/repositories"
)

type mysqlUserRepository struct{}

// NewMysqlUserRepository returns implement of user repository interface
func NewMysqlUserRepository() repositories.UserRepository {
	return &mysqlUserRepository{}
}

// CheckLogin checks if username and password are corrects
func (m *mysqlUserRepository) CheckLogin(username, password string) (models.User, error) {
	var user models.User

	if len(username) == 0 || len(password) == 0 {
		return user, nil
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
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Lastname,
			&user.Firstname,
			&user.CreatedAt,
			&user.DeletedAt)

		lib.CheckError(err, 0)
	}

	return user, err
}

// ChangePassword changes user password in database
func (m *mysqlUserRepository) ChangePassword(user *models.User, password string) bool {
	return len(database.Orm.Model(&user).Update("password", password).GetErrors()) == 0
}

// GetAllSqlRows returns an array of *sql.Rows of all users
func (m *mysqlUserRepository) GetAllSqlRows(limit uint) (rows *sql.Rows, err error) {
	if limit <= 0 {
		query := `
			SELECT id, username, password, lastname, firstname, created_at, updated_at, deleted_at
			FROM users`
		rows, err = database.Select(query)
	} else {
		query := `
			SELECT id, username, password, lastname, firstname, created_at, updated_at, deleted_at
			FROM users
			LIMIT ?`
		rows, err = database.Select(query, limit)
	}

	return
}
