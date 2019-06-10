package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/jinzhu/gorm"
)

// Migration20191006175832_AddRolesTable migration
func Migration20191006175832_AddRolesTable(db *gorm.DB) {
	db.AutoMigrate(&models.Role{})
}
