package database

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

var (
	// DB is the connection handle
	DB *sql.DB
)

// Open opens database connection
func Open() {
	databaseConfig := lib.Config.Database

	db, err := sql.Open(
		databaseConfig.Driver,
		databaseConfig.User+":"+databaseConfig.Password+
			"@tcp("+databaseConfig.Host+":"+strconv.Itoa(databaseConfig.Port)+")"+
			"/"+databaseConfig.Name+"?parseTime=true&loc=Europe%2FParis")
	lib.CheckError(err, 0)

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
	if lib.Config.SQLLog.Level >= 1 {
		elapsed := time.Since(start)
		limit := lib.Config.SQLLog.Limit

		if limit == 0.0 || lib.Config.SQLLog.DisplayOverLimit || elapsed.Seconds() >= limit {
			lib.SQLLog(elapsed, query, args)
		}
	}
}
