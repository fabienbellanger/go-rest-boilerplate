package orm

import (
	migrations2 "github.com/fabienbellanger/go-rest-boilerplate/database/orm/migrations"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/jinzhu/gorm"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE="+lib.Config.Database.Engine)

	migrations2.Migration20190607000000CreateUserTable(db)
	migrations2.Migration20191006175832AddRolesTable(db)
	migrations2.Migration20190803215634CreateApplicationModuleActionTables(db)
	// migrations.Migration20190308215349CreateApplicationModuleActionTables(db)
	// migrations.Migration20190803215634CreateApplicationModuleActionTables(db)
}
