package echo

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/routes/web"
)

// StartEchoServer starts Echo web server
func StartEchoServer(port int) {
	e := initEchoServer()

	// Start server
	// ------------
	s := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}

// initEchoServer initializes Echo server
func initEchoServer() *echo.Echo {
	// La config pour le CORS est-elle correcte ?
	// ------------------------------------------
	if !lib.IsServerConfigCorrect() {
		lib.CheckError(errors.New("no allow origins defined in config file"), 1)
	}

	// Echo instance
	e := echo.New()

	// Startup banner
	// --------------
	if viper.GetString("environment") != "production" {
		e.HideBanner = true
	}

	// Logger
	// ------
	initLogger(e)

	// Recover
	// -------
	e.Use(middleware.Recover())

	// CORS & Secure middlewares
	// -------------------------
	initCorsAndSecurity(e)

	// HTTP errors management
	// ----------------------
	e.HTTPErrorHandler = customHTTPErrorHandler

	// Profilage
	// ---------
	if viper.GetBool("server.pprof") {
		web.NewWebPprofRoute(e.Group("")).PprofRoutes()
	}

	// Liste des routes
	// ----------------
	initRoutes(e)

	// Favicon, static files and template renderer
	// -------------------------------------------
	initStaticFilesAndTemplates(e)

	return e
}
