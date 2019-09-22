package logs

import (
	"regexp"
	"strings"

	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/fabienbellanger/go-rest-boilerplate/repositories"
)

type fileLogsRepository struct{}

// NewfileLogsRepository returns implement of logs repository interface
func NewfileLogsRepository() repositories.LogsRepository {
	return &fileLogsRepository{}
}

// GetAll returns all logs
func (m *fileLogsRepository) GetAll() ([]models.LogFile, error) {
	logs := make([]models.LogFile, 0)

	errorLog := models.LogFile{Error: &models.LogErrorFile{
		Source:    "ERR",
		Timestamp: "2019-09-20 14:17:58",
		Message:   "Internal server error",
	}}

	logs = append(logs, errorLog)

	return logs, nil
}

// parseError returns a variable of LogErrorFile type
func parseError(line string) (log models.LogErrorFile) {
	var regex = regexp.MustCompile(`(ERR)(?:[| \t]+)([\d-: ]{19})(?:[| \t]+)(.*)`)

	found := regex.FindAllStringSubmatch(line, -1)
	if len(found) == 1 {
		for _, match := range found {
			if len(match) == 4 {
				log.Source = strings.Trim(match[1], " ")
				log.Timestamp = strings.Trim(match[2], " ")
				log.Message = strings.Trim(match[3], " ")
			}
		}
	}

	return
}

// parseEcho returns a variable of LogEchoFile type
func parseEcho(line string) (log models.LogEchoFile) {
	var regex = regexp.MustCompile(`(ECHO)(?:[| \t]+)([\d-: ]{19})(?:[| \t]+)([\d]{3})(?:[| \t]+)([0-9a-z.\p{L}]+)(?:[| \t]*)([A-Z]+)(?:[| \t]*)(.*)`)

	found := regex.FindAllStringSubmatch(line, -1)
	if len(found) == 1 {
		for _, match := range found {
			if len(match) == 7 {
				log.Source = strings.Trim(match[1], " ")
				log.Timestamp = strings.Trim(match[2], " ")
				log.Code = strings.Trim(match[3], " ")
				log.Latency = strings.Trim(match[4], " ")
				log.Method = strings.Trim(match[5], " ")
				log.Uri = strings.Trim(match[6], " ")
			}
		}
	}

	return
}

// parseSql returns a variable of LogSqlFile type
func parseSql(line string) (log models.LogSqlFile) {
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
				log.Paramters = strings.Trim(match[7], " ")
			}
		}
	}

	return
}
