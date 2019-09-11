package repositories

import (
	"database/sql"

	"github.com/fabienbellanger/go-rest-boilerplate/models"
)

type UserRepository interface {
	CheckLogin(username, password string) (models.User, error)
	ChangePassword(user *models.User, password string) bool
	GetAllSqlRows(limit uint) (rows *sql.Rows, err error)
}