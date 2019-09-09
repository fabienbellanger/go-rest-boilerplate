package database

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/migrations"
)

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE="+viper.GetString("database.engine"))

	migrations.Migration20190607000000CreateUserTable(db)
}
