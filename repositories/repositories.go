package repositories

import "github.com/fabienbellanger/go-rest-boilerplate/models"

type UserRepository interface {
	CheckLogin(username, password string) (models.User, error)
	ChangePassword(user *models.User, password string) bool
}
