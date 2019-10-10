package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/fabienbellanger/go-rest-boilerplate/websockets"
)

func init() {
	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(WebSocketCommand)
}

// WebSocketCommand : Web command
var WebSocketCommand = &cobra.Command{
	Use:   "websocket",
	Short: "Launch the WebSocket server",

	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|--------------------------|
|                          |
| Launch WebSockets server |
|                          |
|--------------------------|

`)
		port := viper.GetInt("webSocketServer.port")

		// Test du port
		// ------------
		if port < 1000 || port > 10000 {
			lib.CheckError(errors.New("a valid port number must be configured (between 1000 and 10000)"), 1)
		}

		fmt.Print("Listening on port  ")
		color.Green(strconv.Itoa(port) + "\n")

		// Connexion à l'ORM
		// -----------------
		database.OpenORM()
		defer database.Orm.Close()

		fmt.Print("Connection to ORM  ")
		color.Green("✔\n\n")

		// Lancement du serveur websocket
		// ------------------------------
		websockets.ServerStart()
	},
}
