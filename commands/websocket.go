package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/database"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/fabienbellanger/go-rest-boilerplate/websockets"
)

var webSocketPort, webSocketDefaultPort int

func init() {
	// Flag
	// ----
	webSocketDefaultPort = 8889
	WebSocketCommand.Flags().IntVarP(&webSocketPort, "port", "p", webSocketDefaultPort, "listened port")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(WebSocketCommand)
}

// WebSocketCommand : Web command
var WebSocketCommand = &cobra.Command{
	Use:   "websocket",
	Short: "Launch the websocket server",

	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|------------------------------------|
|                                    |
| Lancement du serveur de WebSockets |
|                                    |
|------------------------------------|

`)

		// Test du port
		// ------------
		if webSocketPort == webSocketDefaultPort && viper.GetInt("webSocketServer.port") != 0 {
			// Si on n'a pas spécifié un port dans la commande, on prend celui du fichier de configuration
			// -------------------------------------------------------------------------------------------
			webSocketPort = viper.GetInt("webSocketServer.port")
		}

		if webSocketPort < 1000 || webSocketPort > 10000 {
			webSocketPort = webSocketDefaultPort
		}

		fmt.Print("Listening on port  ")
		color.Green(strconv.Itoa(webSocketPort) + "\n")

		// Connexion à l'ORM
		// -----------------
		database.OpenORM()
		defer database.Orm.Close()

		fmt.Print("Connection to ORM  ")
		color.Green("✔\n\n")

		// Lancement du serveur websocket
		// ------------------------------
		websockets.ServerStart(webSocketPort)
	},
}
