package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/routes/echo"
)

var port, defaultPort int

func init() {
	// Flag
	// ----
	defaultPort = 8888
	APICommand.Flags().IntVarP(&port, "port", "p", defaultPort, "listened port")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(APICommand)
}

// APICommand : API command
var APICommand = &cobra.Command{
	Use:   "api",
	Short: "Launch the web server API",

	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|--------------------------------|
|                                |
| Lancement du serveur Web (API) |
|                                |
|--------------------------------|

`)

		// Test du port
		// ------------
		if port == defaultPort && viper.GetInt("server.port") != 0 {
			// Si on n'a pas spécifié un port dans la commande, on prend celui du fichier de configuration
			// -------------------------------------------------------------------------------------------
			port = viper.GetInt("server.port")
		}

		if port < 1000 || port > 10000 {
			port = defaultPort
		}

		fmt.Print("Listening on port \t")
		color.Green(strconv.Itoa(port) + "\n")

		// Connexion à MySQL
		// -----------------
		if !lib.IsDatabaseConfigCorrect() {
			err := errors.New("no or missing database information in settings file")
			lib.CheckError(err, 1)
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
		echo.StartEchoServer(port)
	},
}
