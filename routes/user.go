package routes

import (
	userController "github.com/fabienbellanger/go-rest-boilerplate/controllers/user"
	"github.com/labstack/echo/v4"
)

// usersRoutes manages users routes
func usersRoutes(g *echo.Group) {
	g.GET("/users", userController.GetUserDetailsHandler)
	g.PATCH("/users/change-password", userController.ChangePassword)
}
