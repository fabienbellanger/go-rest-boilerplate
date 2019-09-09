package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

var rootCommand = &cobra.Command{
	Use:     "Golang Rest API boilerplate",
	Short:   "Golang Rest API boilerplate",
	Long:    "Golang Rest API boilerplate",
	Version: viper.GetString("version"),
}

// Execute starts Cobra
func Execute() {
	// Initialisation de la configuration
	// ----------------------------------
	viper.SetConfigFile("./config.toml")
	err := viper.ReadInConfig()
	lib.CheckError(err, 1)

	lib.InitConfig("config.toml")

	// Lancement de la commande racine
	// -------------------------------
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
