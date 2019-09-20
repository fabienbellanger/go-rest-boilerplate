package logs

import (
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
