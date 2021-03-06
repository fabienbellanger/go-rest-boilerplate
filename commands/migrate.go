package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

var force bool

func init() {
	MigrateCommand.Flags().BoolVar(&force, "force", false, "Force migration")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(MigrateCommand)
}

// MigrateCommand create database migration
var MigrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Make database migrations",
	Long:  "Make database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|---------------------|
|                     |
| Database migrations |
|                     |
|---------------------|

`)

		if force {
			// Connexion à l'ORM
			// -----------------
			fmt.Println("Connecting to GORM...")
			database.OpenORM()
			defer database.Orm.Close()
			lib.DisplaySuccessMessage("Connection OK\n")

			// Migrate the schema
			// TODO: Peut-être long si la base de données contient beaucoup de tables
			fmt.Println("\nStarting migrations...")
			database.Migrate(database.Orm)
			lib.DisplaySuccessMessage("Migrations OK\n")
		} else {
			fmt.Println("Use --force flag to make migrations")
			fmt.Println("")
		}
	},
}
