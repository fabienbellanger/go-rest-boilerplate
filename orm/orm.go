package orm

import (
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Model overrights gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Connect connects to the database and links to the ORM
func Connect() {
	databaseConfig := lib.Config.Database

	db, err := gorm.Open(databaseConfig.Driver,
		databaseConfig.User+":"+databaseConfig.Password+
			"@tcp("+databaseConfig.Host+":"+strconv.Itoa(databaseConfig.Port)+")"+
			"/"+databaseConfig.Name+"?parseTime=true&loc="+databaseConfig.Timezone+
			"&charset="+databaseConfig.Charset)
	lib.CheckError(err, 1)
	defer db.Close()
}
