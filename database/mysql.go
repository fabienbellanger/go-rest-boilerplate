package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

var (
	// DB is the connection handle
	DB *sql.DB
)

// Open opens database connection
func Open() {
	db, err := sql.Open(
		viper.GetString("database.driver"),
		viper.GetString("database.user")+":"+viper.GetString("database.password")+
			"@tcp("+viper.GetString("database.host")+":"+viper.GetString("database.port")+")"+
			"/"+viper.GetString("database.name")+"?parseTime=true&loc="+
			viper.GetString("database.timezone")+
			"&charset="+viper.GetString("database.charset"))
	lib.CheckError(err, 1)

	db.SetMaxOpenConns(viper.GetInt("database.maxOpenConnections"))
	db.SetMaxIdleConns(viper.GetInt("database.maxIdleConnections"))
	db.SetConnMaxLifetime(time.Duration(viper.GetInt("database.maxLifetimeConnection")) * time.Minute)

	DB = db
}

// prepareQuery prepares query
func prepareQuery(query string) *sql.Stmt {
	statement, err := DB.Prepare(query)
	lib.CheckError(err, 0)

	return statement
}

// executeQuery executes request of type INSERT, UPDATE or DELETE
func executeQuery(query string, args ...interface{}) (sql.Result, error) {
	// Start timer
	start := time.Now()

	statement := prepareQuery(query)
	defer statement.Close()

	result, err := statement.Exec(args...)
	lib.CheckError(err, 0)

	// Query log
	logRequest(start, query, args)

	return result, err
}

// Select request
func Select(query string, args ...interface{}) (*sql.Rows, error) {
	// Start timer
	start := time.Now()

	statement := prepareQuery(query)
	defer statement.Close()

	rows, err := statement.Query(args...)
	lib.CheckError(err, 0)

	// Query log
	logRequest(start, query, args)

	return rows, err
}

// Insert request
func Insert(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	lib.CheckError(err, 0)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	lib.CheckError(err, 0)

	if err != nil {
		return 0, err
	}

	return id, err
}

// Update request
func Update(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	lib.CheckError(err, 0)

	if err != nil {
		return 0, err
	}

	affect, err := result.RowsAffected()
	lib.CheckError(err, 0)

	if err != nil {
		return 0, err
	}

	return affect, err
}

// Delete request
func Delete(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	lib.CheckError(err, 0)

	if err != nil {
		return 0, err
	}

	affect, err := result.RowsAffected()
	lib.CheckError(err, 0)

	if err != nil {
		return 0, err
	}

	return affect, err
}

// logRequest writes query log to Gin default writer
func logRequest(start time.Time, query string, args ...interface{}) {
	if viper.GetInt("sql_log.level") >= 1 {
		elapsed := time.Since(start)
		limit := viper.GetFloat64("sql_log.limit")

		if limit == 0.0 || viper.GetBool("sql_log.displayOverLimit") || elapsed.Seconds() >= limit {
			lib.SQLLog(elapsed, query, args)
		}
	}
}
