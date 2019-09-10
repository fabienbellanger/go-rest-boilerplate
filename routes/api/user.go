package api

import (
	"github.com/labstack/echo/v4"

	userHandler "github.com/fabienbellanger/go-rest-boilerplate/handlers/user"
	"github.com/fabienbellanger/go-rest-boilerplate/routes"
)

type apiUserRoute struct {
	Group *echo.Group
}

// NewApiUserRoute returns implement of api user routes
func NewApiUserRoute(g *echo.Group) routes.ApiUserRoutes {
	return &apiUserRoute{
		Group: g,
	}
}

// UsersRoutes manages users routes
func (r *apiUserRoute) UsersRoutes() {
	userHandler := userHandler.NewUserHandler()

	r.Group.GET("/users", userHandler.GetUserDetailsHandler)
	r.Group.PATCH("/users/change-password", userHandler.ChangePassword)
}
