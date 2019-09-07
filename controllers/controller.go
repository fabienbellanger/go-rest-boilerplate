package controllers

import "github.com/labstack/echo/v4"

// UserController type details
type UserController interface {
	GetUserDetailsHandler(c echo.Context) error
	LoginHandler(c echo.Context) error
	ChangePassword(c echo.Context) error
}
