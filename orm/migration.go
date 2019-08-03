package orm

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/migrations"
	"github.com/jinzhu/gorm"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE="+lib.Config.Database.Engine)

	migrations.Migration20190607000000CreateUserTable(db)
	migrations.Migration20191006175832AddRolesTable(db)
	migrations.Migration20190803215634CreateApplicationModuleActionTables(db)
	// migrations.Migration20190308215349CreateApplicationModuleActionTables(db)
	// migrations.Migration20190803215634CreateApplicationModuleActionTables(db)
}
