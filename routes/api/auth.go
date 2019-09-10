package api

import (
	"github.com/labstack/echo/v4"

	userHandler "github.com/fabienbellanger/go-rest-boilerplate/handlers/user"
	"github.com/fabienbellanger/go-rest-boilerplate/routes"
)

type apiAuthRoute struct {
	Group *echo.Group
}

// NewApiAuthRoute returns implement of api authentication routes
func NewApiAuthRoute(g *echo.Group) routes.ApiAuthRoutes {
	return &apiAuthRoute{
		Group: g,
	}
}

// AuthRoutes manages authentication routes for Echo
func (r *apiAuthRoute) AuthRoutes() {
	userHandler := userHandler.NewUserHandler()

	r.Group.POST("/login", userHandler.LoginHandler)
}
