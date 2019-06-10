package commands

import (
	"errors"
	"fmt"

	"code.cloudfoundry.org/bytefmt"
	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var isInit, isDump, dumpDataOnly, dumpStructureOnly bool

func init() {
	DatabaseCommand.Flags().BoolVarP(&isInit, "init", "i", false, "Database initialization")
	DatabaseCommand.Flags().BoolVarP(&isDump, "dump", "d", false, "Database dump structure and data")
	DatabaseCommand.Flags().BoolVarP(&dumpDataOnly, "data-only", "o", false, "Database dump data only")
	DatabaseCommand.Flags().BoolVarP(&dumpStructureOnly, "structure-only", "s", false, "Database dump structure only")

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

		if isInit {
			// Initialisation
			// --------------
			var confirm = "Y"

			fmt.Println("If a database already exists, data will be deleted")
			fmt.Print("Do you really want to initalize database (Y/n): ")
			_, err := fmt.Scanf("%s", &confirm)

			if err != nil || confirm == "n" {
				fmt.Print("\n\n -> Database initialization: ")
				color.Yellow("Operation aborded\n\n")
			} else {
				fmt.Print("\n\n -> Database initialization: ")

				database.InitDatabase()

				color.Green("Success\n\n")
			}
		} else if isDump {
			// Dump
			// ----
			fmt.Print(" -> Database dump: ")

			fileName, fileSize := database.DumpDatabase(dumpStructureOnly, dumpDataOnly)

			color.Green(fileName + " (" + bytefmt.ByteSize(uint64(fileSize)) + ") successfully created\n\n")
		}
	},
}
