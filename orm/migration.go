package orm

import (
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/jinzhu/gorm"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
