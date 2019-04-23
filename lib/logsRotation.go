package lib

import (
	"apiticSellers/server/lib"
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

const NB_FILES_TO_ARCHIVE = 7

// ExecuteLogsRotation launches logs rotation
//
// Description:
// Every day at 00:00
func ExecuteLogsRotation() {
	logFileName = Config.Log.Filename

	_, err := os.OpenFile(logFileName, os.O_RDONLY, 0755)

	if err != nil {
		lib.CheckError(err, -1)
	}

	// Recherche des fichiers de log archivés
	// --------------------------------------
	logFiles, err := findLogFile()

	if err != nil {
		lib.CheckError(err, -2)
	}

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

	// fmt.Println(logFiles)

	// Archivage des NB_FILES_TO_ARCHIVE derniers fichiers
	// ---------------------------------------------------
	// TODO
	/*var buf bytes.Buffer
	archiveWriter := gzip.NewWriter(&buf)

	for _, file := range logFiles {
		archiveWriter.Name = file.name
		archiveWriter.Comment = file.comment
		archiveWriter.ModTime = file.modTime

		if _, err := archiveWriter.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		if err := archiveWriter.Close(); err != nil {
			log.Fatal(err)
		}

		archiveWriter.Reset(&buf)
	}*/

	// os.Exit(0)

	// Déplacement du fichier de log
	// -----------------------------
	err = os.Rename(logFileName, logFileName+".1")

	if err != nil {
		lib.CheckError(err, -3)
	}

	// Création du nouveau fichier logFileName
	// ----------------------------------------
	_, err = os.Create(logFileName)

	if err != nil {
		// Le fichier de log n'existe pas
		// ------------------------------
		lib.CheckError(err, -4)
	}
}

// findLogFile returns the list of log files
func findLogFile() ([]logFile, error) {
	var isLogFile bool
	var lastPoint, fileNameSuffix int
	var fileNameWithoutSuffix string
	var logFiles = make([]logFile, 0)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
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
