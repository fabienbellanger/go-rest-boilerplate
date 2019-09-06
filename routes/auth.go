package routes

import (
	userController "github.com/fabienbellanger/go-rest-boilerplate/controllers/user"
	"github.com/labstack/echo/v4"
)

// authRoutes manages authentication routes for Echo
func authRoutes(g *echo.Group) {
	g.POST("/login", userController.LoginHandler)
}
