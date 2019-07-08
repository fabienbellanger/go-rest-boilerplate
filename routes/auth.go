package routes

import (
	"github.com/fabienbellanger/go-rest-boilerplate/controllers"
	"github.com/labstack/echo/v4"
)

// userLogin is used for binding data in login route
type userLogin struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

// authRoutes manages authentication routes for Echo
func authRoutes(e *echo.Echo, g *echo.Group) {
	// Liste des routes
	g.POST("/login", controllers.LoginHandler)
}
