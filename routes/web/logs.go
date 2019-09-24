package web

import (
	"html/template"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/handlers/web"
	"github.com/fabienbellanger/go-rest-boilerplate/routes"
)

type webLogsRoute struct {
	Group *echo.Group
}

var TemplateFuncMap = template.FuncMap{
	"displayHttpCode": DisplayHttpCode,
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

// DisplayHttpCode displays right color in function of HTTP code
func DisplayHttpCode(code int) string {
	if code >= 200 && code < 300 {
		return "success"
	} else if code >= 500 && code < 600 {
		return "danger"
	}

	return "warning"
}
