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
	accessLogs, _ := h.repository.GetAccessLogs(5)
	errorLogs, _ := h.repository.GetErrorLogs(5)
	// accessLogsJson, _ := json.Marshal(accessLogs)
	// errorLogsJson, _ := json.Marshal(errorLogs)

	return c.Render(http.StatusOK, "logs/index.gohtml", map[string]interface{}{
		"title":      "Logs server interface",
		"accessLogs": accessLogs,
		"errorLogs":  errorLogs,
	})
}
