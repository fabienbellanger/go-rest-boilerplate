package orm

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/migrations"
	"github.com/jinzhu/gorm"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE="+lib.Config.Database.Engine)

	migrations.Migration20190607000000_CreateUserTable(db)
}
