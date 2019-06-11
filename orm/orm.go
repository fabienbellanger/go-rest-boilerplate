package orm

import (
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB is the connection handle
var DB *gorm.DB

// Open connects to the database and links to the ORM
func Open() {
	databaseConfig := lib.Config.Database

	db, err := gorm.Open(databaseConfig.Driver,
		databaseConfig.User+":"+databaseConfig.Password+
			"@tcp("+databaseConfig.Host+":"+strconv.Itoa(databaseConfig.Port)+")"+
			"/"+databaseConfig.Name+"?parseTime=true&loc="+databaseConfig.Timezone+
			"&charset="+databaseConfig.Charset)
	lib.CheckError(err, 1)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(lib.Config.ORM.MaxIdleConnections)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(lib.Config.ORM.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Duration(lib.Config.ORM.MaxLifetimeConnection) * time.Minute)

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(false)

	// Enable Logger, show detailed log
	db.LogMode(lib.Config.ORM.EnabledLogs)

	// Migrate the schema
	// TODO: Pluôt mettre dans une commande ? Peut-être long si la base de données contient beaucoup de tables
	// Migrate(db)

	DB = db
}
