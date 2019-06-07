package commands

import (
	"errors"
	"fmt"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm"
	"github.com/fabienbellanger/go-rest-boilerplate/routes"
	"github.com/spf13/cobra"
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
		fmt.Println(`

|--------------------------------|
|                                |
| Lancement du serveur Web (API) |
|                                |
|--------------------------------|`)

		// Test du port
		// ------------
		if port == defaultPort && lib.Config.Server.Port != 0 {
			// Si on n'a pas spécifié un port dans la commande, on prend celui du fichier de configuration
			// -------------------------------------------------------------------------------------------
			port = lib.Config.Server.Port
		}

		if port < 1000 || port > 10000 {
			port = defaultPort
		}

		// Connexion à MySQL
		// -----------------
		if !lib.IsDatabaseConfigCorrect() {
			err := errors.New("no or missing database information in settings file")
			lib.CheckError(err, 1)
		}

		database.Open()
		defer database.DB.Close()

		// Connexion à l'ORM
		// -----------------
		orm.Open()
		defer orm.DB.Close()

		// Lancement du serveur web
		// ------------------------
		routes.StartServer(port)
	},
}
