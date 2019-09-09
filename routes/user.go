package routes

import (
	"github.com/labstack/echo/v4"

	userHandler "github.com/fabienbellanger/go-rest-boilerplate/handlers/user"
)

// usersRoutes manages users routes
func usersRoutes(g *echo.Group) {
	userHandler := userHandler.NewUserHandler()

	g.GET("/users", userHandler.GetUserDetailsHandler)
	g.PATCH("/users/change-password", userHandler.ChangePassword)
}
