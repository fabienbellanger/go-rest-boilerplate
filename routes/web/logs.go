package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

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
	r.Group.GET("/logs", func(c echo.Context) error {
		return c.Render(http.StatusOK, "logs/index.gohtml", map[string]interface{}{
			"title": "Logs interface",
		})
	})
}
