package commands

import (
	"fmt"

	"github.com/fabienbellanger/go-rest-boilerplate/routes"
	"github.com/spf13/cobra"
)

var port, defaultPort int

func init() {
	// Flag
	// ----
	defaultPort = 8080
	WebCommand.Flags().IntVarP(&port, "port", "p", defaultPort, "listened port")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(WebCommand)
}

// WebCommand : Web command
var WebCommand = &cobra.Command{
	Use:   "web",
	Short: "Launch the web server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`

|--------------------------|
|                          |
| Lancement du serveur Web |
|                          |
|--------------------------|`)

		// Test du port
		// ------------
		if port < 1000 || port > 10000 {
			port = defaultPort
		}

		// Lancement du serveur web
		// ------------------------
		routes.StartServer(port)
	},
}
