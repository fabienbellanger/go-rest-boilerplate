package orm

import (
	"github.com/fabienbellanger/go-rest-boilerplate/database/migrations"
	"github.com/jinzhu/gorm"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	migrations.Migration_20190607_1(db)
}
