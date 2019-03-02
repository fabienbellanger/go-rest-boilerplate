package commands

import (
	"errors"
	"fmt"
	"github.com/cloudfoundry/bytefmt"
	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Init : Initialisation de la base de données
var Init bool

// Dump : Dump de la base de données
var Dump bool

func init() {
	DatabaseCommand.Flags().BoolVarP(&Init, "init", "i", false, "Database initialization")
	DatabaseCommand.Flags().BoolVarP(&Dump, "dump", "d", false, "Database dump")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(DatabaseCommand)
}

// DatabaseCommand : Database command
var DatabaseCommand = &cobra.Command{
	Use:   "db",
	Short: "Database operations",
	Long:  "Database operations: initialisation and dump",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|---------------------|
|                     |
| Database operations |
|                     |
|---------------------|

		`)

		// Connexion à MySQL
		// -----------------
		if !lib.IsDatabaseConfigCorrect() {
			err := errors.New("no or missing database information in settings file")
			lib.CheckError(err, 1)
		}

		database.Open()
		defer database.DB.Close()

		if Init {
			// Initialisation
			// --------------
			var confirm = "Y"

			fmt.Println("If a database already exists, data will be deleted")
			fmt.Print("Do you really want to initalize database (Y/n): ")
			fmt.Scanf("%s", &confirm)
			// _, err := fmt.Scanf("%s", &confirm)
			// lib.CheckError(err, 0)

			if confirm == "Y" {
				fmt.Print("\n\n -> Database initialization:\t")

				database.InitDatabase()

				color.Green("Success\n\n")
			} else {
				fmt.Print("\n\n -> Database initialization:\t")
				color.Yellow("Operation aborded\n\n")
			}
		} else if Dump {
			// Dump
			// ----
			fmt.Print(" -> Database dump: ")

			fileName, fileSize := database.DumpDatabase()

			color.Green(fileName + " (" + bytefmt.ByteSize(uint64(fileSize)) + ") successfully created\n\n")
		}
	},
}
