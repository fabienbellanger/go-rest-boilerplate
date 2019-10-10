package database

import (
	"os"
	"os/exec"
	"time"

	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// InitDatabase : Initialisation de la base de données
func InitDatabase() {
	// Requêtes
	// --------
	queries := make([]string, 0)

	// User
	queries = append(queries, "DROP TABLE IF EXISTS user")
	queries = append(queries, `
		CREATE TABLE users (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			username varchar(191) NOT NULL,
			password varchar(128) NOT NULL,
			lastname varchar(100) NOT NULL,
			firstname varchar(100) NOT NULL,
			created_at datetime NOT NULL,
			deleted_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	queries = append(queries, "CREATE UNIQUE INDEX index_users_username ON users(username)")
	queries = append(queries, "CREATE INDEX index_users_password ON users(password)")
	queries = append(queries, "CREATE INDEX index_users_deleted_at ON users(deleted_at)")

	transaction, err := DB.Begin()
	lib.CheckError(err, 1)
	defer func() {
		// Rollback the transaction after the function returns.
		// If the transaction was already committed, this will do nothing.
		_ = transaction.Rollback()
	}()

	for _, query := range queries {
		// Execute the query in the transaction.
		_, err := transaction.Exec(query)
		lib.CheckError(err, 1)
	}

	// Commit the transaction.
	err = transaction.Commit()
	lib.CheckError(err, 1)
}

// DumpDatabase : Dump de la base de données
func DumpDatabase(structureOnly bool, dataOnly bool) (string, int) {
	// Exécution de la commande
	// ------------------------
	var dumpCommand *exec.Cmd

	if !structureOnly && dataOnly {
		dumpCommand = exec.Command("mysqldump",
			"-u"+viper.GetString("database.user"),
			"-p"+viper.GetString("database.password"),
			"--no-create-info",
			"--single-transaction",
			viper.GetString("database.name"))
	} else if structureOnly && !dataOnly {
		dumpCommand = exec.Command("mysqldump",
			"-u"+viper.GetString("database.user"),
			"-p"+viper.GetString("database.password"),
			"--no-data",
			"--single-transaction",
			viper.GetString("database.name"))
	} else {
		dumpCommand = exec.Command("mysqldump",
			"-u"+viper.GetString("database.user"),
			"-p"+viper.GetString("database.password"),
			"--single-transaction",
			viper.GetString("database.name"))
	}
	dumpCommand.Dir = "."
	dumpOutput, err := dumpCommand.Output()
	lib.CheckError(err, 1)

	// Création du fichier
	// -------------------
	dumpFileName := "dump_" + viper.GetString("database.name") + "_" + time.Now().Format("2006-01-02_150405") + ".sql"
	file, err := os.Create(dumpFileName)
	lib.CheckError(err, 2)
	defer file.Close()

	size, err := file.Write(dumpOutput)
	lib.CheckError(err, 3)
	return dumpFileName, size
}
