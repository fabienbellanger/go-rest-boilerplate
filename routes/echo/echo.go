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

type Host struct {
	Echo *echo.Echo
}

// StartServer starts Echo web server
func StartServer(port int) {
	// Hosts
	hosts := map[string]*Host{}
	subdomain := ""
	domain := viper.GetString("server.domain") + ":" + strconv.Itoa(port)

	// Initialisation du serveur d'API
	// -------------------------------
	api := initApiServer()
	subdomain = viper.GetString("server.apiSubDomain")
	if subdomain != "" {
		subdomain += "."
	}
	hosts[subdomain+domain] = &Host{api}

	// Initialisation du serveur pour le client (Web)
	// ----------------------------------------------
	client := initClientServer()
	subdomain = viper.GetString("server.clientSubDomain")
	if subdomain != "" {
		subdomain += "."
	}
	hosts[subdomain+domain] = &Host{client}

	// Start server
	// ------------
	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		request := c.Request()
		response := c.Response()

		host := hosts[request.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(response, request)
		}

		return
	})

	// Startup banner
	// --------------
	if viper.GetString("environment") == "production" {
		e.HideBanner = true
	}

	s := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}

// initApiServer initializes API server
func initApiServer() *echo.Echo {
	// La config pour le CORS est-elle correcte ?
	// ------------------------------------------
	if !lib.IsServerConfigCorrect() {
		lib.CheckError(errors.New("no allow origins defined in config file"), 1)
	}

	// Echo instance
	e := echo.New()

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

	// Liste des routes
	// ----------------
	initApiRoutes(e)

	// Favicon, static files and template renderer
	// -------------------------------------------
	// TODO
	initStaticFilesAndTemplates(e)

	return e
}

// initClientServer initializes API for a client
func initClientServer() *echo.Echo {
	// Echo instance
	e := echo.New()

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
	initClientRoutes(e)

	// Favicon, static files and template renderer
	// -------------------------------------------
	// TODO
	initStaticFilesAndTemplates(e)

	return e
}
