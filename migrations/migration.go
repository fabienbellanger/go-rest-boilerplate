package migrations

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/jinzhu/gorm"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE="+lib.Config.Database.Engine)

	Migration20190607000000CreateUserTable(db)
	Migration20191006175832AddRolesTable(db)
	Migration20190803215634CreateApplicationModuleActionTables(db)
	// migrations.Migration20190308215349CreateApplicationModuleActionTables(db)
	// migrations.Migration20190803215634CreateApplicationModuleActionTables(db)
}
