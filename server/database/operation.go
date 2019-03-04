package database

import (
	"apiticSellers/server/lib"
	"os"
	"os/exec"
	"time"
)

// InitDatabase : Initialisation de la base de données
func InitDatabase() {
	// Requètes
	// --------
	queries := make([]string, 0)

	// User
	// ----
	queries = append(queries, "DROP TABLE IF EXISTS user")
	queries = append(queries, `
		CREATE TABLE user (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			username varchar(191) NOT NULL,
			password varchar(128) NOT NULL,
			lastname varchar(100) NOT NULL,
			firstname varchar(100) NOT NULL,
			created_at datetime NOT NULL,
			deleted_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	queries = append(queries, "CREATE UNIQUE INDEX idx_user_username ON user(username)")
	queries = append(queries, "CREATE INDEX idx_user_password ON user(password)")
	queries = append(queries, "CREATE INDEX idx_user_deleted_at ON user(deleted_at)")

	// Seller
	// ------
	queries = append(queries, "DROP TABLE IF EXISTS seller")
	queries = append(queries, `
		CREATE TABLE seller (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			lastname varchar(100) NOT NULL,
			firstname varchar(100) NOT NULL,
			email varchar(191),
			phone varchar(16),
			comment text,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			deleted_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	queries = append(queries, "CREATE INDEX idx_seller_deleted_at ON seller(deleted_at)")

	// Device
	// ------
	queries = append(queries, "DROP TABLE IF EXISTS device")
	queries = append(queries, `
		CREATE TABLE device (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(100) NOT NULL,
			serial_number varchar(100) NOT NULL,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			deleted_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	queries = append(queries, "CREATE INDEX idx_device_deleted_at ON device(deleted_at)")

	// Event
	// -----
	queries = append(queries, "DROP TABLE IF EXISTS event")
	queries = append(queries, `
		CREATE TABLE event (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			device_id int(11) unsigned NOT NULL,
			data text,
			created_at datetime NOT NULL,
			updated_at datetime NOT NULL,
			deleted_at datetime NULL DEFAULT NULL,
			PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`)
	queries = append(queries, `
		ALTER TABLE event ADD CONSTRAINT fk_event_device_id
		FOREIGN KEY (device_id REFERENCES device(id)
		ON DELETE CASCADE`)
	queries = append(queries, "CREATE INDEX idx_event_deleted_at ON event(deleted_at)")

	transaction, err := DB.Begin()
	lib.CheckError(err, 1)

	defer func() {
		// Rollback the transaction after the function returns.
		// If the transaction was already commited, this will do nothing.
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
func DumpDatabase() (string, int) {
	// Exécution de la commande
	// ------------------------
	dumpCommand := exec.Command("mysqldump",
		"-u"+lib.Config.Database.User,
		"-p"+lib.Config.Database.Password,
		lib.Config.Database.Name)
	dumpCommand.Dir = "."
	dumpOutput, err := dumpCommand.Output()
	lib.CheckError(err, 1)

	// Création du fichier
	// -------------------
	dumpFileName := "dump_" + lib.Config.Database.Name + "_" + time.Now().Format("2006-01-02_150405") + ".sql"
	file, err := os.Create(dumpFileName)
	lib.CheckError(err, 2)

	defer file.Close()

	size, err := file.Write(dumpOutput)
	lib.CheckError(err, 3)

	return dumpFileName, size
}
