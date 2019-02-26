package commands

import (
	"os"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:     "Golang Rest API boilerplate",
	Short:   "Golang Rest API boilerplate",
	Long:    "Golang Rest API boilerplate",
	Version: lib.Config.Version,
}

// Execute starts Cobra
func Execute() {
	// Initialisation de la configuration
	// ----------------------------------
	lib.InitConfig()

	// Lancement de la commande racine
	// -------------------------------
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
