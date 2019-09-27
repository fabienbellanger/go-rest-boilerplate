package commands

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var logsErrors bool
var logsSql bool
var logsAccess bool
var logsAll bool
var logsCsvSeparator string

func init() {
	LogsExportCommand.Flags().BoolVarP(&logsErrors, "errors", "e", false, "Export errors logs")
	LogsExportCommand.Flags().BoolVarP(&logsAccess, "access", "a", false, "Export access logs")
	LogsExportCommand.Flags().BoolVarP(&logsSql, "sql", "s", false, "Export SQL logs")
	LogsExportCommand.Flags().BoolVarP(&logsAll, "all", "A", false, "Export all logs")
	LogsExportCommand.Flags().StringVarP(&logsCsvSeparator, "separator", "c", ";", "CSV separator (',', ';', 'tab')")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(LogsExportCommand)
}

// LogsExportCommand exports logs in CSV files
var LogsExportCommand = &cobra.Command{
	Use:   "logs-export",
	Short: "Export logs in CSV files",
	Long:  "Export logs in CSV files",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|-------------|
|             |
| Logs export |
|             |
|-------------|

`)

		logs := make([]string, 0, 3)
		if logsAll {
			logs = append(logs, "access")
			logs = append(logs, "errors")
			logs = append(logs, "sql")
		} else {
			if logsAccess {
				logs = append(logs, "access")
			}

			if logsErrors {
				logs = append(logs, "errors")
			}

			if logsSql {
				logs = append(logs, "sql")
			}
		}

		if len(logs) == 0 {
			logs = append(logs, "access")
			logs = append(logs, "errors")
			logs = append(logs, "sql")
		}

		if logsCsvSeparator != "," && logsCsvSeparator != ";" && logsCsvSeparator != "tab" {
			// TODO: error
		}

		// Traitement
		// ----------
		exportLogs(logs)
	},
}

// exportLogs exports selected logs
func exportLogs(logs []string) {
	log.Printf("Separator: %s\nLogs: %+v\n", logsCsvSeparator, logs)
}
