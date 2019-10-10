package repositories

import (
	"database/sql"

	"github.com/fabienbellanger/go-rest-boilerplate/models"
)

// UserRepository lists methods
type UserRepository interface {
	CheckLogin(username, password string) (models.User, error)
	ChangePassword(user *models.User, password string) bool
	GetAllSqlRows(limit uint) (rows *sql.Rows, err error)
}

// LogsRepository lists methods
type LogsRepository interface {
	GetAccessLogs(size int) ([]models.LogFile, error)
	GetErrorLogs(size int) ([]models.LogFile, error)
	GetSqlLogs(size int) ([]models.LogFile, error)
}
