package database

import (
	"database/sql"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"strconv"
	"time"
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
			"@tcp("+databaseConfig.Host+":"+strconv.Itoa(databaseConfig.Port)+")"+
			"/"+databaseConfig.Name+"?parseTime=true&loc=Europe%2FParis")
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
	var start time.Time

	if lib.Config.SqlLog.Level >= 1 {
		start = time.Now()
	}

	statement := prepareQuery(query)
	defer statement.Close()

	result, err := statement.Exec(args...)
	lib.CheckError(err, 0)

	if lib.Config.SqlLog.Level >= 1 {
		// TODO: Gérer la variable limit du fichier de configuration
		elapsed := time.Since(start)
		lib.SqlLog(elapsed, query)
	}

	return result, err
}

// Select : Exécution d'une requête
func Select(query string, args ...interface{}) (*sql.Rows, error) {
	var start time.Time

	if lib.Config.SqlLog.Level >= 1 {
		start = time.Now()
	}

	statement := prepareQuery(query)
	defer statement.Close()

	rows, err := statement.Query(args...)
	lib.CheckError(err, 0)

	if lib.Config.SqlLog.Level >= 1 {
		// TODO: Gérer la variable limit du fichier de configuration
		elapsed := time.Since(start)
		lib.SqlLog(elapsed, query, args)
	}

	return rows, err
}

// Insert : Requête d'insertion
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

// Update : Requête de mise à jour
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

// Delete : Requête de suppression
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
