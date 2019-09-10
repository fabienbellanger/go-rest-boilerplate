package handlers

import "github.com/labstack/echo/v4"

// UserHandler type details
type UserHandler interface {
	GetUserDetailsHandler(c echo.Context) error
	LoginHandler(c echo.Context) error
	ChangePassword(c echo.Context) error
}
