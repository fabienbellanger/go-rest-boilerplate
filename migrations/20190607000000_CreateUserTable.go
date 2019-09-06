package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/jinzhu/gorm"
)

// Migration20190607000000CreateUserTable migration
func Migration20190607000000CreateUserTable(db *gorm.DB) {
	db.AutoMigrate(&models.User{})

	db.Model(&models.User{}).AddIndex("idx_created_at", "created_at")
	db.Model(&models.User{}).AddIndex("idx_deleted_at", "deleted_at")
}
