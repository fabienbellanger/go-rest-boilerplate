package user

import (
	"crypto/sha512"
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
		println()
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
