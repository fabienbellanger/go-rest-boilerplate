package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/jinzhu/gorm"
)

// Migration20190607_1 migration
func Migration20190607_1(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
