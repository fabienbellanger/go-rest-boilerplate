package commands

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/routes/echo"
)

func init() {
	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(ServerCommand)
}

// ServerCommand : Server command
var ServerCommand = &cobra.Command{
	Use:   "serve",
	Short: "Launch the Web server",

	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|-------------------|
|                   |
| Launch Web server |
|                   |
|-------------------|

`)
		// Informations
		// ------------
		fmt.Print("Version\t\t\t")
		color.Green(viper.GetString("version") + "\n")
		fmt.Print("Environment\t\t")
		color.Green(viper.GetString("environment") + "\n")

		// Test du port
		// ------------
		port := viper.GetInt("server.port")
		if port < 10 || port > 10000 {
			lib.CheckError(errors.New("a valid port number must be configured (between 1000 and 10000)"), 1)
		}

		// Connexion à MySQL
		// -----------------
		if !lib.IsDatabaseConfigCorrect() {
			err := errors.New("no or missing database information in settings file")
			lib.CheckError(err, 2)
		}

		database.Open()
		defer database.DB.Close()

		fmt.Print("Connection to database \t")
		color.Green("✔\n")

		// Connexion à l'ORM
		// -----------------
		database.OpenORM()
		defer database.Orm.Close()

		fmt.Print("Connection to ORM \t")
		color.Green("✔\n\n")

		// Lancement du serveur web
		// ------------------------
		echo.StartServer()
	},
}
