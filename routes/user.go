package routes

import (
	"github.com/fabienbellanger/go-rest-boilerplate/controllers"
	"github.com/labstack/echo/v4"
)

// usersRoutes manages users routes
func usersRoutes(e *echo.Echo, g *echo.Group) {
	g.GET("/users", controllers.GetUserDetailsHandler)
	g.PATCH("/users/change-password", controllers.ChangePassword)
}
