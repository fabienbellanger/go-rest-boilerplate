package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/jinzhu/gorm"
)

// Migration20191006175832AddRolesTable migration
func Migration20191006175832AddRolesTable(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})

	db.Model(&models.Role{}).AddForeignKey("parent_id", "roles(id)", "CASCADE", "CASCADE")

	db.Model(&models.Role{}).AddIndex("idx_created_at", "created_at")
	db.Model(&models.Role{}).AddIndex("idx_deleted_at", "deleted_at")
}
