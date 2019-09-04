package routes

import (
	"github.com/fabienbellanger/go-rest-boilerplate/controllers"
	"github.com/labstack/echo/v4"
)

// authRoutes manages authentication routes for Echo
func authRoutes(e *echo.Echo, g *echo.Group) {
	g.POST("/login", controllers.LoginHandler)
}
