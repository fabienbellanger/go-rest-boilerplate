package api

import (
	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/handlers/api"
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
	userHandler := api.NewUserHandler()

	r.Group.POST("/login", userHandler.LoginHandler)
}
