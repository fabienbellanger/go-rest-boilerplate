package web

import (
	"encoding/json"
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

// GetLogs returns all logs
func (h *LogsHandler) GetLogs(c echo.Context) error {
	logs, _ := h.repository.GetAll()
	logsJson, _ := json.Marshal(logs)

	return c.Render(http.StatusOK, "logs/index.gohtml", map[string]interface{}{
		"title": "Logs interface",
		"logs":  string(logsJson),
	})
}
