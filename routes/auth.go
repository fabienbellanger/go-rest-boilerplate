package routes

import (
	"github.com/labstack/echo/v4"

	userHandler "github.com/fabienbellanger/go-rest-boilerplate/handlers/user"
)

// authRoutes manages authentication routes for Echo
func authRoutes(g *echo.Group) {
	userHandler := userHandler.NewUserHandler()

	g.POST("/login", userHandler.LoginHandler)
}
