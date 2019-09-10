package migrations

import (
	"github.com/jinzhu/gorm"

	"github.com/fabienbellanger/go-rest-boilerplate/models"
)

// Migration20190803215634CreateApplicationModuleActionTables migration
func Migration20190803215634CreateApplicationModuleActionTables(db *gorm.DB) {
	db.AutoMigrate(&models.Application{})
	db.AutoMigrate(&models.Module{})
	db.AutoMigrate(&models.Action{})
}
