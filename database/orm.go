package database

import (
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// Database is the connection handle
var Orm *gorm.DB

// Open connects to the database and links to the ORM
func OpenORM() {
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

	// var user models.User

	// for i := 0; i < 100000; i++ {
	// 	user = models.User{
	// 		Username:  "ffgfgfghhfghfhgfgfhgfghfghfhgfhgfh" + strconv.Itoa(i),
	// 		Password:  "gjgjghjgjhgjhghjfrserhkhjhklljjkbhjvftxersgdghjjkhkljkbhftd",
	// 		Lastname:  "njuftydfhgjkjlkjlkjlkhjkhu",
	// 		Firstname: "jkggkjkl,,lm,kljkvgf"}

	// 	db.NewRecord(user)
	// 	db.Create(&user)
	// }

	Orm = db
}
