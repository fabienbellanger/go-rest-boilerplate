package routes

import (
	"github.com/labstack/echo/v4"

	userController "github.com/fabienbellanger/go-rest-boilerplate/controllers/user"
)

// usersRoutes manages users routes
func usersRoutes(g *echo.Group) {
	g.GET("/users", userController.GetUserDetailsHandler)
	g.PATCH("/users/change-password", userController.ChangePassword)
}
