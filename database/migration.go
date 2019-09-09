package database

import (
	"github.com/jinzhu/gorm"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/migrations"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE="+lib.Config.Database.Engine)

	migrations.Migration20190607000000CreateUserTable(db)
}
