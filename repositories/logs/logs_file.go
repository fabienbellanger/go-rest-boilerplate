package logs

import (
	"fmt"
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

	parseError(`ERR  | 2019-09-20 21:51:55 | [template: no template "logs/index.gohtml" associated with template "index.gohtml"]`)
	parseEcho(`ECHO | 2019-09-20 21:51:55 |  500 | 326.921µs	| GET	| /logs`)

	logs = append(logs, errorLog)

	return logs, nil
}

// parseError returns a variable of LogErrorFile type
func parseError(err string) (log models.LogErrorFile) {
	var regex = regexp.MustCompile(`(ERR ) | ([\d-: ]{19}) | (.*)`)

	found := regex.FindAllString(err, -1)
	if len(found) == 3 {
		for i, match := range found {
			if i == 0 {
				log.Source = strings.Trim(match, " ")
			} else if i == 1 {
				log.Timestamp = strings.Trim(match, " ")
			} else {
				log.Message = strings.Trim(match, " ")
			}
		}
	}

	fmt.Printf("%+v\n", log)

	return
}

// parseEcho returns a variable of LogEchoFile type
func parseEcho(err string) (log models.LogEchoFile) {
	var regex = regexp.MustCompile(`(ECHO) | ([\d-: ]{19}) |  ([\d]{3}) | ([0-9a-z.\p{L}]+)| ([A-Z]+)| (.*)`)

	found := regex.FindAllString(err, -1)
	if len(found) == 6 {
		for i, match := range found {
			if i == 0 {
				log.Source = strings.Trim(match, " ")
			} else if i == 1 {
				log.Timestamp = strings.Trim(match, " ")
			} else if i == 2 {
				log.Code = strings.Trim(match, " ")
			} else if i == 3 {
				log.Latency = strings.Trim(match, " ")
			} else if i == 4 {
				log.Method = strings.Trim(match, " ")
			} else {
				log.Uri = strings.Trim(match, " ")
			}
		}
	}

	fmt.Printf("%+v\n", log)

	return
}
