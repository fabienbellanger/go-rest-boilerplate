package logs

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	logs := getFileLines("logs/access.log", 5)

	return logs, nil
}

// GetErrorLogs returns error logs
func (m *fileLogsRepository) GetErrorLogs(size int) ([]models.LogFile, error) {
	logs := getFileLines("logs/error.log", 5)

	return logs, nil
}

// GetSqlLogs returns SQL logs
func (m *fileLogsRepository) GetSqlLogs(size int) ([]models.LogFile, error) {
	logs := getFileLines("logs/sql.log", 5)

	return logs, nil
}

// Récupère les lignes du fichier
func getFileLines(fileName string, size int) []models.LogFile {
	file, err := os.Open(fileName)
	defer file.Close()
	lib.CheckError(err, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []models.LogFile
	if size != 0 {
		lines = make([]models.LogFile, 0, size)
		i := 0
		for scanner.Scan() && i < size {
			text := scanner.Text()
			if strings.Contains(fileName, viper.GetString("log.server.accessFilename")) {
				// Logs d'accès
				// ------------
				if line, isFound := parseEcho(text); isFound {
					lines = append(lines, models.LogFile{Echo: &line})

					i++
				}
			} else if strings.Contains(fileName, viper.GetString("log.server.errorFilename")) {
				// Logs d'erreur
				// -------------
				if line, isFound := parseError(text); isFound {
					lines = append(lines, models.LogFile{Error: &line})

					i++
				}
			} else if strings.Contains(fileName, viper.GetString("log.sql.sqlFilename")) {
				// Logs SQL
				// --------
				if line, isFound := parseSql(text); isFound {
					lines = append(lines, models.LogFile{Sql: &line})

					i++
				}
			}
		}
	} else {
		// Si size == 0, alors on prend toutes les lignes
		lines = make([]models.LogFile, 0)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.Contains(fileName, viper.GetString("log.server.accessFilename")) {
				// Logs d'accès
				// ------------
				if line, isFound := parseEcho(text); isFound {
					lines = append(lines, models.LogFile{Echo: &line})
				}
			} else if strings.Contains(fileName, viper.GetString("log.server.errorFilename")) {
				// Logs d'erreur
				// -------------
				if line, isFound := parseError(text); isFound {
					lines = append(lines, models.LogFile{Error: &line})
				}
			} else if strings.Contains(fileName, viper.GetString("log.sql.sqlFilename")) {
				// Logs SQL
				// --------
				if line, isFound := parseSql(text); isFound {
					lines = append(lines, models.LogFile{Sql: &line})
				}
			}
		}
	}

	return lines
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
