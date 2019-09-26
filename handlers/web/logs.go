package web

import (
	"fmt"
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
	sqlLogs, _ := h.repository.GetSqlLogs(5)

	return c.Render(http.StatusOK, "logs/index.gohtml", map[string]interface{}{
		"title":      "Server Logs Interface",
		"accessLogs": accessLogs,
		"errorLogs":  errorLogs,
		"sqlLogs":    sqlLogs,
	})
}

// LoadLogs returns logs
func (h *LogsHandler) LoadLogs(c echo.Context) error {
	query := c.Request().URL.Query()
	delete(query, "_")
	fmt.Printf("%+v\n", query)

	return c.String(http.StatusOK, "OK")
}
