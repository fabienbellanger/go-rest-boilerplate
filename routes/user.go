package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/fabienbellanger/go-rest-boilerplate/controllers"
)

// usersRoutes manages users routes
func usersRoutes(g *echo.Group) {
	g.GET("/users", controllers.GetUserDetailsHandler)
	g.PATCH("/users/change-password", controllers.ChangePassword)
}
