package commands

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	// LogCommand.Flags().BoolVarP(&Init, "init", "i", false, "Database initialization")
	// LogCommand.Flags().BoolVarP(&Dump, "dump", "d", false, "Database dump")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(LogCommand)
}

// LogCommand : Logs rotation command
var LogCommand = &cobra.Command{
	Use:   "log",
	Short: "Gin logs rotation",
	Long:  "Gin logs rotation",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|--------------------|
|                    |
| Gin logs rotation  |
|                    |
|--------------------|

		`)

		// Launches logs rotation
		lib.ExecuteLogsRotation()
	},
}
