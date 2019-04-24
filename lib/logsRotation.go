package lib

import (
	"apiticSellers/server/lib"
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type logFile struct {
	name     string
	suffix   int
	fullname string
}

var logFileName string

const LOGS_PATH = "logs/"

// ExecuteLogsRotation launches logs rotation
//
// Description:
// Every day at 00:00
func ExecuteLogsRotation() {
	logFileName = LOGS_PATH + Config.Log.FileName

	logFile, err := os.OpenFile(logFileName, os.O_RDONLY, 0755)

	if err != nil {
		lib.CheckError(err, -1)
	}

	defer logFile.Close()

	// Recherche des fichiers de log archivés
	// --------------------------------------
	logFiles, err := findLogFile()

	if err != nil {
		lib.CheckError(err, -2)
	}

	// Rotation des fichiers
	// ---------------------
	makeLogsRotation(logFiles)

	// Archivage des fichiers
	// ----------------------
	err = makeLogsArchiving(logFiles)

	if err != nil {
		lib.CheckError(err, 0)
	}

	// Déplacement du fichier de log
	// -----------------------------
	err = os.Rename(logFileName, logFileName+".1")

	if err != nil {
		lib.CheckError(err, -3)
	}

	// Création du nouveau fichier logFileName
	// ---------------------------------------
	logFile, err = os.Create(logFileName)

	if err != nil {
		// Le fichier de log n'existe pas
		// ------------------------------
		lib.CheckError(err, -4)
	}

	defer logFile.Close()
}

// findLogFile returns the list of log files
func findLogFile() ([]logFile, error) {
	var isLogFile bool
	var lastPoint, fileNameSuffix int
	var fileNameWithoutSuffix string
	var logFiles = make([]logFile, 0)

	err := filepath.Walk(LOGS_PATH, func(path string, info os.FileInfo, err error) error {
		isLogFile, _ = regexp.Match(`^`+logFileName+`.[\d]+$`, []byte(path))

		if isLogFile {
			// On récupère les fichiers de log archivés uniquement
			// ---------------------------------------------------
			lastPoint = strings.LastIndex(path, ".")

			if lastPoint != -1 {
				fileNameWithoutSuffix = path[:lastPoint]
				fileNameSuffix, err = strconv.Atoi(path[lastPoint+1:])

				if err == nil && fileNameWithoutSuffix == logFileName {
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

	return logFiles, err
}

// findArchiveName returns the name of the next archive file
func findArchiveName() (string, error) {
	var fileSuffix int
	fileName := logFileName
	regex := regexp.MustCompile(`^` + logFileName + `.([\d]+).zip$`)
	maxFileSuffix := 0

	err := filepath.Walk(LOGS_PATH, func(path string, info os.FileInfo, err error) error {
		regexResult := regex.FindAllSubmatch([]byte(path), -1)

		for _, matchMessage := range regexResult {
			if len(matchMessage) == 2 {
				fileSuffix, _ = strconv.Atoi(string(matchMessage[1]))

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
func makeLogsArchiving(logFiles []logFile) error {
	// Nombre de fichiers à archiver
	// -----------------------------
	nbFilesToArchive := Config.Log.NbFilesToArchive

	if nbFilesToArchive <= 0 {
		nbFilesToArchive = 1
	}

	if len(logFiles) < nbFilesToArchive {
		return nil
	}

	// 1. Recherche du nom de la prochaine archive
	// -------------------------------------------
	archiveFileName, err := findArchiveName()

	if err != nil {
		return err
	}

	// 2. Création de l'archive
	// ------------------------
	newZipFile, err := os.Create(archiveFileName)

	if err != nil {
		return err
	}

	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// 3. Ajout des fichiers de log à l'archive
	// ----------------------------------------

	// 3.1. Conversion en tableau de nom de fichier
	// --------------------------------------------
	var logFilesName = make([]string, 0)

	for _, file := range logFiles {
		logFilesName = append(logFilesName, file.fullname)
	}

	// 3.2. Ajout à l'archive
	// ----------------------
	for _, file := range logFilesName {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}

	// 4. Supression des fichiers logs archivés
	// ----------------------------------------
	for _, file := range logFilesName {
		err = os.Remove(file)

		if err != nil {
			lib.CheckError(err, 0)
		}
	}

	return nil
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

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
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
