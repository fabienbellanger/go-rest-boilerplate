package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/models/action"
	"github.com/fabienbellanger/go-rest-boilerplate/models/application"
	"github.com/fabienbellanger/go-rest-boilerplate/models/module"
	"github.com/jinzhu/gorm"
)

// Migration20190803215634CreateApplicationModuleActionTables migration
func Migration20190803215634CreateApplicationModuleActionTables(db *gorm.DB) {
	db.AutoMigrate(&application.Application{})
	db.AutoMigrate(&module.Module{})
	db.AutoMigrate(&action.Action{})
}
