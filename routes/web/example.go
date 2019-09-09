package web

import (
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/routes"

	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

type webExampleRoute struct {
	Group *echo.Group
}

// NewWebExampleRoute returns implement of web example routes
func NewWebExampleRoute(g *echo.Group) routes.WebExampleRoutes {
	return &webExampleRoute{
		Group: g,
	}
}

// ExampleRoutes lists example routes
func (r *webExampleRoute) ExampleRoutes() {
	// Test page for websockets
	r.Group.GET("/websockets", func(c echo.Context) error {
		return c.Render(http.StatusOK, "example/websockets.gohtml", map[string]interface{}{
			"title":        "Websockets example",
			"webSocketUrl": strconv.Itoa(lib.Config.WebSocketServer.Port),
		})
	})

	// Test page for VueJS
	r.Group.GET("/vuejs", func(c echo.Context) error {
		return c.Render(http.StatusOK, "example/vuejs.gohtml", map[string]interface{}{
			"title": "VueJS example",
		})
	})
}
