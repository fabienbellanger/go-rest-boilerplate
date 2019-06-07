package commands

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const migrationsPath = "./database/migrations/"

var migrationFileName string

func init() {
	MigratationCommand.Flags().StringVarP(&migrationFileName, "name", "n", "", "Migration file name")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(MigratationCommand)
}

// MigratationCommand create database migration
var MigratationCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration",
	Long:  "Database migration",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|--------------------|
|                    |
| Database migration |
|                    |
|--------------------|

		`)

		timePrefix := time.Now().Format("20060201150405")
		migrationFileNamePath := migrationsPath + timePrefix + "_" + migrationFileName + ".go"

		// Le répertoire existe t-il ?
		// ---------------------------
		if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
			lib.CheckError(errors.New(migrationsPath+" directory does not exist"), -1)
		}

		// Le fichier existe t-il ?
		// ------------------------
		if _, err := os.Stat(migrationFileNamePath); err == nil {
			lib.CheckError(errors.New(migrationFileNamePath+" file already exists"), -2)
		}

		// Création du fichier
		// -------------------
		file, err := os.Create(migrationFileNamePath)
		lib.CheckError(err, -3)
		defer file.Close()

		// Ecriture dans le fichier
		// ------------------------
		content := []byte(`package migrations

import "github.com/jinzhu/gorm"

//  Migration` + timePrefix + "_" + migrationFileName + ` migration
func Migration` + timePrefix + "_" + migrationFileName + `(db *gorm.DB) {
	
}
`)
		_, err = file.Write(content)
		if err != nil {
			// Suppression du fichier
			err := os.Remove(migrationFileNamePath)
			lib.CheckError(err, -5)
		}
		lib.CheckError(err, -4)

		color.New(color.FgGreen).Print("[✔️] ")
		fmt.Println("File " + timePrefix + "_" + migrationFileName + ".go created\n")
	},
}
