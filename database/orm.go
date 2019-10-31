package database

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL dialect
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// Orm represents database is the connection handle
var Orm *gorm.DB

// OpenORM connects to the database and links to the ORM
func OpenORM() {
	db, err := gorm.Open(
		viper.GetString("database.driver"),
		viper.GetString("database.user")+":"+viper.GetString("database.password")+
			"@tcp("+viper.GetString("database.host")+":"+viper.GetString("database.port")+")"+
			"/"+viper.GetString("database.name")+"?parseTime=true&loc=GMT"+ //viper.GetString("database.timezone")+
			"&charset="+viper.GetString("database.charset"))
	lib.CheckError(err, 1)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(viper.GetInt("orm.maxIdleConnections"))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(viper.GetInt("orm.maxOpenConnections"))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("orm.maxLifetimeConnection")) * time.Minute)

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(false)

	// Enable Logger, show detailed log
	db.LogMode(viper.GetBool("orm.enabledLogs"))

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
