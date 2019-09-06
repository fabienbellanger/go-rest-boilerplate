package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/models/role"
	"github.com/jinzhu/gorm"
)

// Migration20191006175832AddRolesTable migration
func Migration20191006175832AddRolesTable(db *gorm.DB) {
	db.AutoMigrate(&role.Role{})

	db.Model(&role.Role{}).AddForeignKey("parent_id", "roles(id)", "CASCADE", "CASCADE")

	db.Model(&role.Role{}).AddIndex("idx_created_at", "created_at")
	db.Model(&role.Role{}).AddIndex("idx_deleted_at", "deleted_at")
}
