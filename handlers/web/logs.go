package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/repositories"
	logsRepository "github.com/fabienbellanger/go-rest-boilerplate/repositories/logs"
)

// LogsHandler
type LogsHandler struct {
	repository repositories.LogsRepository
}

// NewLogsHandler
func NewLogsHandler() *LogsHandler {
	return &LogsHandler{
		repository: logsRepository.NewfileLogsRepository(),
	}
}

// GetLogs returns logs
func (h *LogsHandler) GetLogs(c echo.Context) error {
	linesNumberMax := 500

	accessLogs, _ := h.repository.GetAccessLogs(linesNumberMax)
	errorLogs, _ := h.repository.GetErrorLogs(linesNumberMax)
	sqlLogs, _ := h.repository.GetSqlLogs(linesNumberMax)

	return c.Render(http.StatusOK, "logs/index.gohtml", map[string]interface{}{
		"title":      "Server Logs Interface",
		"accessLogs": accessLogs,
		"errorLogs":  errorLogs,
		"sqlLogs":    sqlLogs,
	})
}
