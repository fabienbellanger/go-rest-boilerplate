package web

import (
	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/handlers/web"
	"github.com/fabienbellanger/go-rest-boilerplate/routes"
)

type webLogsRoute struct {
	Group *echo.Group
}

// NewWebLogsRoute returns implement of web logs interface routes
func NewWebLogsRoute(g *echo.Group) routes.WebLogsRoutes {
	return &webLogsRoute{
		Group: g,
	}
}

// LogsRoutes lists logs routes
func (r *webLogsRoute) LogsRoutes() {
	logsHandler := web.NewLogsHandler()

	r.Group.GET("/logs", logsHandler.GetLogs)
}
