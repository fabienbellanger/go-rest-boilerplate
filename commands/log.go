package commands

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

type logFile struct {
	name     string
	suffix   int
	fullname string
}

var logErrorFileName string
var logAccessFileName string

func init() {
	// Ajout de la commande à la commande racine
	rootCommand.AddCommand(LogCommand)
}

// LogCommand : Logs rotation command
var LogCommand = &cobra.Command{
	Use:   "log",
	Short: "Server logs rotation",
	Long:  "Server logs rotation",
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow(`

|----------------------|
|                      |
| Server logs rotation |
|                      |
|----------------------|

`)

		// Launches logs rotation
		executeLogsRotation()
	},
}

// executeLogsRotation launches logs rotation
func executeLogsRotation() {
	// Le répertoire existe t-il ?
	// ---------------------------
	if _, err := os.Stat(viper.GetString("log.server.dirPath")); os.IsNotExist(err) {
		lib.CheckError(errors.New(viper.GetString("log.server.dirPath")+" directory does not exist"), 1)
	}

	// Vérifie que les fichiers d'erreur et d'accès existent
	// -----------------------------------------------------
	checkLogFiles()

	// Recherche des anciens fichiers de log non archivés
	// --------------------------------------------------
	logErrorFiles := findLogFile(logErrorFileName)
	logAccessFiles := findLogFile(logAccessFileName)

	// Rotation des fichiers de log non archivés
	// -----------------------------------------
	makeLogsRotation(logErrorFiles)
	makeLogsRotation(logAccessFiles)

	// Archivage des fichiers de log
	// -----------------------------
	makeLogsArchiving(logErrorFiles, logErrorFileName)
	makeLogsArchiving(logAccessFiles, logAccessFileName)

	// Déplace le contenu du fichier courant dans le fichier préfixé par .1
	// et remet à zéro le contenu du fichier courant
	// --------------------------------------------------------------------
	createNewLogFile(logErrorFileName)
	createNewLogFile(logAccessFileName)

	lib.DisplaySuccessMessage("Logs rotation successfully completed\n")
}

// createNewLogFile moves current file content in file prefix by .1 and reset current file content
func createNewLogFile(fileName string) {
	copyFile(fileName, fileName+".tmp")

	// Déplacement des fichiers de log
	// -------------------------------
	err := os.Rename(fileName+".tmp", fileName+".1")
	lib.CheckError(err, 1)

	// Remise à zéro des fichiers courants
	// -----------------------------------
	err = ioutil.WriteFile(fileName, []byte(""), 0644)
	lib.CheckError(err, 2)
}

// copyFile copies a source file to a destination file from their path names
func copyFile(sourceName, destinationName string) {
	sourceFileStat, err := os.Stat(sourceName)
	lib.CheckError(err, 1)

	if !sourceFileStat.Mode().IsRegular() {
		lib.CheckError(fmt.Errorf("%s is not a regular file", sourceName), 2)
	}

	source, err := os.Open(sourceName)
	lib.CheckError(err, 3)
	defer source.Close()

	destination, err := os.Create(destinationName)
	lib.CheckError(err, 4)
	defer destination.Close()

	_, err = io.Copy(destination, source)
	lib.CheckError(err, 5)
}

// checkLogFiles checks if log files exists
func checkLogFiles() {
	logErrorFileName = viper.GetString("log.server.dirPath") + viper.GetString("log.server.errorFilename")

	_, err := os.OpenFile(logErrorFileName, os.O_RDWR, 0755)
	if err != nil {
		lib.CheckError(errors.New("log file "+viper.GetString("log.server.errorFilename")+" does not exists"), 2)
	}

	logAccessFileName = viper.GetString("log.server.dirPath") + viper.GetString("log.server.accessFilename")

	_, err = os.OpenFile(logErrorFileName, os.O_RDWR, 0755)
	if err != nil {
		lib.CheckError(errors.New("log file "+viper.GetString("log.server.errorFilename")+" does not exists"), 2)
	}
}

// findLogFile returns the list of log files
func findLogFile(logFilename string) []logFile {
	logFiles := make([]logFile, 0)

	// On parcours tous les fichiers du dossier
	// ----------------------------------------
	err := filepath.Walk(viper.GetString("log.server.dirPath"), func(path string, info os.FileInfo, err error) error {
		isLogFile, _ := regexp.Match(`^`+logFilename+`.[\d]+$`, []byte(path))

		if isLogFile {
			// On récupère les fichiers de log archivés uniquement
			// ---------------------------------------------------
			lastPoint := strings.LastIndex(path, ".")

			if lastPoint != -1 {
				fileNameWithoutSuffix := path[:lastPoint]
				fileNameSuffix, err := strconv.Atoi(path[lastPoint+1:])

				if err == nil && fileNameWithoutSuffix == logFilename {
					// Ajout du fichier à la liste des fichiers de logs archivés
					// ---------------------------------------------------------
					logFiles = append(logFiles, logFile{
						fileNameWithoutSuffix,
						fileNameSuffix,
						fileNameWithoutSuffix + "." + strconv.Itoa(fileNameSuffix),
					})
				}
			}
		}

		return nil
	})

	lib.CheckError(err, 3)

	return logFiles
}

// findArchiveName returns the name of the next archive file
func findArchiveName(fileName string) (string, error) {
	regex := regexp.MustCompile(`^` + fileName + `.([\d]+).zip$`)
	maxFileSuffix := 0

	// On parcours tous les fichiers du dossier
	// ----------------------------------------
	err := filepath.Walk(viper.GetString("log.server.dirPath"), func(path string, info os.FileInfo, err error) error {
		regexResult := regex.FindAllSubmatch([]byte(path), -1)

		for _, matchMessage := range regexResult {
			if len(matchMessage) == 2 {
				fileSuffix, _ := strconv.Atoi(string(matchMessage[1]))

				if fileSuffix > maxFileSuffix {
					maxFileSuffix = fileSuffix
				}
			}
		}

		return nil
	})

	// Nom de la future archive
	// ------------------------
	fileName = fileName + "." + strconv.Itoa(maxFileSuffix+1) + ".zip"

	return fileName, err
}

// makeLogsRotation makes log rotation by renamming files
func makeLogsRotation(logFiles []logFile) {
	// Décalage des fichiers d'archivage
	// ---------------------------------
	for index := range logFiles {
		logFiles[index].suffix++
		logFiles[index].fullname = logFiles[index].name + "." + strconv.Itoa(logFiles[index].suffix)
	}

	// Renommage des fichiers d'archivage
	// ----------------------------------
	for i := len(logFiles) - 1; i >= 0; i-- {
		_ = os.Rename(logFiles[i].name+"."+strconv.Itoa(logFiles[i].suffix-1), logFiles[i].fullname)
	}
}

// makeLogsArchiving makes logs archiving
func makeLogsArchiving(logFiles []logFile, logFilename string) {
	// Nombre de fichiers à archiver
	// -----------------------------
	nbFilesToArchive := viper.GetInt("log.server.nbFilesToArchive")

	if nbFilesToArchive <= 0 {
		nbFilesToArchive = 1
	}

	if len(logFiles) < nbFilesToArchive {
		return
	}

	// Recherche du nom de la prochaine archive
	// ----------------------------------------
	archiveFileName, err := findArchiveName(logFilename)
	lib.CheckError(err, 0)

	// Création de l'archive
	// ---------------------
	newZipFile, err := os.Create(archiveFileName)
	lib.CheckError(err, 0)
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Ajout des fichiers de log à l'archive
	// -------------------------------------

	// Conversion en tableau de nom de fichier
	// ---------------------------------------
	var logFilesName = make([]string, 0)

	for _, file := range logFiles {
		logFilesName = append(logFilesName, file.fullname)
	}

	// Ajout à l'archive
	// -----------------
	for _, file := range logFilesName {
		if err = addFileToZip(zipWriter, file); err != nil {
			lib.CheckError(err, 0)
		}
	}

	// Supression des fichiers logs archivés
	// -------------------------------------
	for _, file := range logFilesName {
		err = os.Remove(file)
		lib.CheckError(err, 0)
	}
}

// addFileToZip adds log file in file archive
func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	// ------------------------
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	// Get the file header
	// -------------------
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = filename

	// Change to deflate to gain better compression
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)

	return err
}
