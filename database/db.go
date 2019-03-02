package database

import (
	"database/sql"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"strconv"
)

var (
	// DB is the connection handle
	DB *sql.DB
)

// Open : Ouverture de la connexion
func Open() {
	databaseConfig := lib.Config.Database

	db, err := sql.Open(
		databaseConfig.Driver,
		databaseConfig.User+":"+databaseConfig.Password+
			"@tcp("+databaseConfig.Host+":"+strconv.Itoa(databaseConfig.Port)+
			")/"+databaseConfig.Name+"?parseTime=true")
	lib.CheckError(err, 0)

	DB = db
}

// prepareQuery : Préparation de la requête
func prepareQuery(query string) *sql.Stmt {
	statement, err := DB.Prepare(query)
	lib.CheckError(err, 0)

	return statement
}

// executeQuery : Exécute une requête de type INSERT, UPDATE ou DELETE
func executeQuery(query string, args ...interface{}) (sql.Result, error) {
	statement := prepareQuery(query)
	defer statement.Close()

	result, err := statement.Exec(args...)
	lib.CheckError(err, 0)

	return result, err
}

// Select : Exécution d'une requête
func Select(query string, args ...interface{}) (*sql.Rows, error) {
	statement := prepareQuery(query)
	defer statement.Close()

	rows, err := statement.Query(args...)
	lib.CheckError(err, 0)

	return rows, err
}

// Insert : Requête d'insertion
func Insert(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	lib.CheckError(err, 0)

	id, err := result.LastInsertId()
	lib.CheckError(err, 0)

	return id, err
}

// Update : Requête de mise à jour
func Update(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	lib.CheckError(err, 0)

	affect, err := result.RowsAffected()
	lib.CheckError(err, 0)

	return affect, err
}

// Delete : Requête de suppression
func Delete(query string, args ...interface{}) (int64, error) {
	result, err := executeQuery(query, args...)
	lib.CheckError(err, 0)

	affect, err := result.RowsAffected()
	lib.CheckError(err, 0)

	return affect, err
}
