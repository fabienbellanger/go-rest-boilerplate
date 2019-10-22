package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	logsRepository "github.com/fabienbellanger/go-rest-boilerplate/repositories/logs"
)

var logsErrors bool
var logsSQL bool
var logsAccess bool
var logsAll bool
var logsCsvSeparator string

func init() {
	LogsExportCommand.Flags().BoolVarP(&logsErrors, "errors", "e", false, "Export errors logs")
	LogsExportCommand.Flags().BoolVarP(&logsAccess, "access", "a", false, "Export access logs")
	LogsExportCommand.Flags().BoolVarP(&logsSQL, "sql", "s", false, "Export SQL logs")
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

			if logsSQL {
				logs = append(logs, "sql")
			}
		}

		if len(logs) == 0 {
			logs = append(logs, "access")
			logs = append(logs, "errors")
			logs = append(logs, "sql")

			logsAll = true
		}

		if logsCsvSeparator != "," && logsCsvSeparator != ";" && logsCsvSeparator != "tab" {
			// TODO: error
		}

		if logsCsvSeparator == "tab" {
			logsCsvSeparator = "\t"
		}

		// Traitement
		// ----------
		exportLogs(logs)
	},
}

// exportLogs exports selected logs
func exportLogs(logs []string) {
	if logsAll {
		createFile(viper.GetString("log.server.errorFilename"))
		createFile(viper.GetString("log.server.accessFilename"))
		createFile(viper.GetString("log.sql.sqlFilename"))
	} else {
		if logsErrors {
			createFile(viper.GetString("log.server.errorFilename"))
		}

		if logsAccess {
			createFile(viper.GetString("log.server.accessFilename"))
		}

		if logsSQL {
			createFile(viper.GetString("log.sql.sqlFilename"))
		}
	}
	fmt.Println("")
}

// createFile creates CSV file and display result to the console
func createFile(fileName string) {
	exportedFileName, err := logsRepository.GetCsvFromFilename(fileName, logsCsvSeparator)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Printf("[%s]  %s file not created\n", red("x"), exportedFileName)
	} else {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Printf("[%s]  %s file created\n", green("✔"), exportedFileName)
	}
}
