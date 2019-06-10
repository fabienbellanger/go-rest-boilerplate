package commands

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const migrationsPath = "./orm/migrations/"

var migrationFileName string

func init() {
	MigratationCommand.Flags().StringVarP(&migrationFileName, "name", "n", "", "Migration file name")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(MigratationCommand)
}

// MigratationCommand create database migration
var MigratationCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Database migration",
	Long:  "Database migration",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|--------------------|
|                    |
| Database migration |
|                    |
|--------------------|

`)

		timePrefix := time.Now().Format("20060201150405")
		migrationFileNamePath := migrationsPath + timePrefix + "_" + migrationFileName + ".go"
		functionName := "Migration" + timePrefix + "_" + migrationFileName

		// Création du fichier de migration
		// --------------------------------
		createMigrationFile(migrationsPath, migrationFileNamePath, functionName, timePrefix)

		// Ajout de la fonction à fin du fichier orm/migration.go
		// ------------------------------------------------------
		updateMigrationsFile(functionName)
	},
}

// createMigrationFile creates migration file
func createMigrationFile(migrationsPath string, migrationFileNamePath string, functionName string, timePrefix string) {
	// Le répertoire existe t-il ?
	// ---------------------------
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		lib.CheckError(errors.New(migrationsPath+" directory does not exist"), -1)
	}

	// Le fichier existe t-il ?
	// ------------------------
	if _, err := os.Stat(migrationFileNamePath); err == nil {
		lib.CheckError(errors.New(migrationFileNamePath+" file already exists"), -2)
	}

	// Création du fichier
	// -------------------
	file, err := os.Create(migrationFileNamePath)
	lib.CheckError(err, -3)
	defer file.Close()

	// Ecriture dans le fichier
	// ------------------------
	content := []byte(`package migrations

import "github.com/jinzhu/gorm"

// ` + functionName + ` migration
func ` + functionName + `(db *gorm.DB) {
	
}
`)
	_, err = file.Write(content)
	if err != nil {
		// Suppression du fichier
		err := os.Remove(migrationFileNamePath)
		lib.CheckError(err, -5)
	}
	lib.CheckError(err, -4)

	lib.DisplaySuccessMessage("File " + timePrefix + "_" + migrationFileName + ".go created\n")
}

// updateMigrationsFile inserts line into migrations file with the commented function
func updateMigrationsFile(functionName string) {
	migrationsFile, err := os.Open("orm/migration.go")
	lib.CheckError(err, -6)
	defer migrationsFile.Close()

	var lines []string
	scanner := bufio.NewScanner(migrationsFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fileContent := ""
	fileIndex := len(lines) - 1
	for i, line := range lines {
		if i == fileIndex {
			fileContent += "	// migrations." + functionName + "(db)\n"
		}
		fileContent += line + "\n"
	}

	err = ioutil.WriteFile("orm/migration.go", []byte(fileContent), 0644)
	lib.CheckError(err, -7)

	lib.DisplaySuccessMessage("File orm/migration.go updated (just uncomment to apply)\n")
}
