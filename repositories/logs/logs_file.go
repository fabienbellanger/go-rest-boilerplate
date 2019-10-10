package logs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hpcloud/tail"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/fabienbellanger/go-rest-boilerplate/repositories"
)

type fileLogsRepository struct{}

// NewfileLogsRepository returns implement of logs repository interface
func NewfileLogsRepository() repositories.LogsRepository {
	return &fileLogsRepository{}
}

// GetAccessLogs returns access logs
func (m *fileLogsRepository) GetAccessLogs(size int) ([]models.LogFile, error) {
	return getFileLines("logs/access.log", size)
}

// GetErrorLogs returns error logs
func (m *fileLogsRepository) GetErrorLogs(size int) ([]models.LogFile, error) {
	return getFileLines("logs/error.log", size)
}

// GetSqlLogs returns SQL logs
func (m *fileLogsRepository) GetSqlLogs(size int) ([]models.LogFile, error) {
	return getFileLines("logs/sql.log", size)
}

// GetCsvFromFilename creates a CSV file of log type
func GetCsvFromFilename(fileName string, sep string) (string, error) {
	// Source
	// ------
	file, err := os.Open("logs/" + fileName)
	if err != nil {
		lib.CheckError(err, 0)
		return "", nil
	}
	defer file.Close()

	// Destination
	// -----------
	var exportFileName string
	if strings.Contains(fileName, viper.GetString("log.server.errorFilename")) {
		exportFileName = "errors"
	} else if strings.Contains(fileName, viper.GetString("log.sql.sqlFilename")) {
		exportFileName = "sql"
	} else if strings.Contains(fileName, viper.GetString("log.server.accessFilename")) {
		exportFileName = "access"
	}
	exportFileName += "_" + time.Now().Format("20060102150405") + ".csv"
	exportFile, err := os.Create(exportFileName)
	if err != nil {
		lib.CheckError(err, 0)
		return "", nil
	}
	defer exportFile.Close()

	// Ecriture du fichier
	// -------------------
	if strings.Contains(fileName, viper.GetString("log.server.errorFilename")) {
		fmt.Fprintf(exportFile, "\"%s\"%s\"%s\"\n", "Timestamp", sep, "Message")
	} else if strings.Contains(fileName, viper.GetString("log.sql.sqlFilename")) {
		fmt.Fprintf(exportFile, "\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"\n",
			"Timestamp", sep,
			"Request", sep,
			"Latency", sep,
			"Query", sep,
			"Parameters")
	} else if strings.Contains(fileName, viper.GetString("log.server.accessFilename")) {
		fmt.Fprintf(exportFile, "\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"\n",
			"Timestamp", sep,
			"Code", sep,
			"Latency", sep,
			"Method", sep,
			"URI")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(fileName, viper.GetString("log.server.errorFilename")) {
			// Logs d'erreur
			// -------------
			if lineParsed, isFound := parseError(line); isFound {
				fmt.Fprintf(exportFile, "\"%s\"%s\"%s\"\n", lineParsed.Timestamp, sep, lineParsed.Message)
			}
		} else if strings.Contains(fileName, viper.GetString("log.sql.sqlFilename")) {
			// Logs SQL
			// --------
			if lineParsed, isFound := parseSql(line); isFound {
				fmt.Fprintf(exportFile, "\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"%s\"\n",
					lineParsed.Timestamp, sep,
					lineParsed.Request, sep,
					lineParsed.Latency, sep,
					lineParsed.Query, sep,
					lineParsed.Parameters)
			}
		} else if strings.Contains(fileName, viper.GetString("log.server.accessFilename")) {
			// Logs d'accès
			// ------------
			if lineParsed, isFound := parseEcho(line); isFound {
				fmt.Fprintf(exportFile, "\"%s\"%s%d%s\"%s\"%s\"%s\"%s\"%s\"\n",
					lineParsed.Timestamp, sep,
					lineParsed.Code, sep,
					lineParsed.Latency, sep,
					lineParsed.Method, sep,
					lineParsed.Uri)
			}
		}
	}

	return exportFileName, nil
}

// Récupère les dernières lignes du fichier
func getFileLines(fileName string, size int) ([]models.LogFile, error) {
	var lines []models.LogFile

	if size < 0 {
		return lines, nil
	} else if size > 500 {
		size = 500
	}

	// Récupération des lignes en partant de la fin du fichier
	// -------------------------------------------------------
	fileLines, err := tail.TailFile(fileName, tail.Config{Follow: false})
	if err != nil {
		lib.CheckError(err, 0)
		return lines, err
	}

	// Traitement des lignes
	// ---------------------
	i := 0
	lines = make([]models.LogFile, 0)
	for fileLine := range fileLines.Lines {
		if i >= size {
			break
		}

		if strings.Contains(fileName, viper.GetString("log.server.accessFilename")) {
			// Logs d'accès
			// ------------
			if line, isFound := parseEcho(fileLine.Text); isFound {
				lines = append(lines, models.LogFile{Echo: &line})
			}
		} else if strings.Contains(fileName, viper.GetString("log.server.errorFilename")) {
			// Logs d'erreur
			// -------------
			if line, isFound := parseError(fileLine.Text); isFound {
				lines = append(lines, models.LogFile{Error: &line})
			}
		} else if strings.Contains(fileName, viper.GetString("log.sql.sqlFilename")) {
			// Logs SQL
			// --------
			if line, isFound := parseSql(fileLine.Text); isFound {
				lines = append(lines, models.LogFile{Sql: &line})
			}
		}

		i++
	}

	return lines, nil
}

// parseError returns a variable of LogErrorFile type
func parseError(line string) (log models.LogErrorFile, isFound bool) {
	var regex = regexp.MustCompile(`(ERR)(?:[| \t]+)([\d-: ]{19})(?:[| \t]+)(.*)`)

	found := regex.FindAllStringSubmatch(line, -1)
	if len(found) == 1 {
		for _, match := range found {
			if len(match) == 4 {
				log.Source = strings.Trim(match[1], " ")
				log.Timestamp = strings.Trim(match[2], " ")
				log.Message = strings.Trim(match[3], " ")

				isFound = true
			}
		}
	}

	return
}

// parseEcho returns a variable of LogEchoFile type
func parseEcho(line string) (log models.LogEchoFile, isFound bool) {
	var regex = regexp.MustCompile(`(ECHO)(?:[| \t]+)([\d-: ]{19})(?:[| \t]+)([\d]{3})(?:[| \t]+)([0-9a-z.\p{L}]+)(?:[| \t]*)([A-Z]+)(?:[| \t]*)(.*)`)

	found := regex.FindAllStringSubmatch(line, -1)
	if len(found) == 1 {
		for _, match := range found {
			if len(match) == 7 {
				code, _ := strconv.Atoi(strings.Trim(match[3], " "))

				log.Source = strings.Trim(match[1], " ")
				log.Timestamp = strings.Trim(match[2], " ")
				log.Code = code
				log.Latency = strings.Trim(match[4], " ")
				log.Method = strings.Trim(match[5], " ")
				log.Uri = strings.Trim(match[6], " ")

				isFound = true
			}
		}
	}

	return
}

// parseSql returns a variable of LogSqlFile type
func parseSql(line string) (log models.LogSqlFile, isFound bool) {
	var regex = regexp.MustCompile(`(SQL)(?:[| \t]+)([\d-: \t]{19})(?:[| \t]+)([A-Z]{3})(?:[| \t]+)([0-9a-z.\p{L}]+)(?:[| \t]+)((.*) (?:[|\t]+)(.*)|)`)

	found := regex.FindAllStringSubmatch(line, -1)
	if len(found) == 1 {
		for _, match := range found {
			if len(match) == 8 {
				log.Source = strings.Trim(match[1], " ")
				log.Timestamp = strings.Trim(match[2], " ")
				log.Request = strings.Trim(match[3], " ")
				log.Latency = strings.Trim(match[4], " ")
				log.Query = strings.Trim(match[6], " ")
				log.Parameters = strings.Trim(match[7], " ")

				isFound = true
			}
		}
	}

	return
}
