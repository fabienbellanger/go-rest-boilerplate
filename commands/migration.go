package commands

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

const migrationsPath = "./migrations/"

var migrationFileName string

func init() {
	MigrationCommand.Flags().StringVarP(&migrationFileName, "name", "n", "", "Migration file name")

	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(MigrationCommand)
}

// MigrationCommand create database migration
var MigrationCommand = &cobra.Command{
	Use:   "make-migration",
	Short: "Create a database migration",
	Long:  "Create a database migration",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|-----------------------------|
|                             |
| Database migration creation |
|                             |
|-----------------------------|

`)

		if migrationFileName == "" {
			lib.CheckError(errors.New("migration file name cannot be empty"), 1)
		}

		timePrefix := time.Now().Format("20060102150405")
		migrationFileNamePath := migrationsPath + timePrefix + "_" + migrationFileName + ".go"
		functionName := "Migration" + timePrefix + migrationFileName

		// Création du fichier de migration
		// --------------------------------
		err := createMigrationFile(migrationsPath, migrationFileNamePath, functionName, timePrefix)
		lib.CheckError(err, 1)

		// Ajout de la fonction à fin du fichier orm/migration.go
		// ------------------------------------------------------
		err = updateMigrationsFile(functionName)
		lib.CheckError(err, 1)
	},
}

// createMigrationFile creates migration file
func createMigrationFile(migrationsPath string, migrationFileNamePath string, functionName string, timePrefix string) error {
	// Le répertoire existe t-il ?
	// ---------------------------
	if _, err := os.Stat(migrationsPath); os.IsNotExist(err) {
		return errors.New(migrationsPath + " directory does not exist")
	}

	// Le fichier existe t-il ?
	// ------------------------
	if _, err := os.Stat(migrationFileNamePath); err == nil {
		return errors.New(migrationFileNamePath + " file already exists")
	}

	// Création du fichier
	// -------------------
	file, err := os.Create(migrationFileNamePath)
	defer file.Close()
	if err != nil {
		return err
	}

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

		return err
	}

	lib.DisplaySuccessMessage("File " + timePrefix + migrationFileName + ".go created\n")

	return nil
}

// updateMigrationsFile inserts line into migrations file with the commented function
func updateMigrationsFile(functionName string) error {
	migrationsFile, err := os.Open("./database/migration.go")
	defer migrationsFile.Close()
	if err != nil {
		return err
	}

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

	err = ioutil.WriteFile("./database/migration.go", []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	lib.DisplaySuccessMessage("File database/migration.go updated (just uncomment to apply)\n")

	return nil
}
