package echo

import (
	"github.com/labstack/echo/v4"
)

// initApiStaticFilesAndTemplates initializes static files and template renderer
func initApiStaticFilesAndTemplates(e *echo.Echo) {
	// Favicon
	// -------
	e.File("/favicon.ico", "assets/favicon.ico")
}

// initWebStaticFilesAndTemplates initializes static files and template renderer
func initClientStaticFilesAndTemplates(e *echo.Echo) {
	// Favicon
	// -------
	e.File("/favicon.ico", "templates/client/favicon.ico")

	// Client index.html
	// -----------------
	// On redirige toutes les pages vers le fichier index.html
	e.File("*", "templates/client/index.html")

	// Assets
	// ------
	e.Static("/js", "./templates/client/js")
	e.Static("/css", "./templates/client/css")
	e.Static("/img", "./templates/client/img")
	e.Static("/fonts", "./templates/client/fonts")
}

// initWebStaticFilesAndTemplates initializes static files and template renderer
func initWebStaticFilesAndTemplates(e *echo.Echo) {
	// Favicon
	// -------
	e.File("/favicon.ico", "assets/favicon.ico")

	// Assets
	// ------
	e.Static("/js", "./assets/js")
	e.Static("/css", "./assets/css")
	e.Static("/img", "./assets/img")
	e.Static("/fonts", "./assets/fonts")

	// Templates
	// ---------
	initTemplates(e)
}
