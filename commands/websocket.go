package commands

import (
	"fmt"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/fabienbellanger/go-rest-boilerplate/websockets"
	"github.com/spf13/cobra"
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
		fmt.Println(`

|--------------------------------|
|                                |
| Lancement du serveur WebSocket |
|                                |
|--------------------------------|`)

		// Test du port
		// ------------
		if webSocketPort == webSocketDefaultPort && lib.Config.WebSocketServer.Port != 0 {
			// Si on n'a pas spécifié un port dans la commande, on prend celui du fichier de configuration
			// -------------------------------------------------------------------------------------------
			webSocketPort = lib.Config.WebSocketServer.Port
		}

		if webSocketPort < 1000 || webSocketPort > 10000 {
			webSocketPort = webSocketDefaultPort
		}

		orm.Open()
		defer orm.DB.Close()

		var user models.User
		orm.DB.First(&user)
		fmt.Printf("%#v\n", user.Lastname)
		user = models.User{}
		orm.DB.First(&user, 2)
		fmt.Printf("%#v\n", user)

		// Lancement du serveur websocket
		// ------------------------------
		websockets.ServerStart(webSocketPort)
	},
}
