package echo

import (
	"errors"
	"net/http"
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
func StartServer() {
	port := viper.GetString("server.port")

	// Sous domaines
	// -------------
	hosts := initSubDomains(port)

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
		Addr:         ":" + port,
		ReadTimeout:  time.Duration(viper.GetInt("server.readTimeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("server.writeTimeout")) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(s))
}

// initSubDomains initializes sub domains
func initSubDomains(port string) map[string]*Host {
	// Hosts
	hosts := map[string]*Host{}
	subdomain := ""
	domain := viper.GetString("server.domain") + ":" + port

	// Initialisation du serveur d'API
	// -------------------------------
	api := initApiServer()
	subdomain = viper.GetString("server.apiSubDomain")
	if subdomain != "" {
		subdomain += "."
	}
	hosts[subdomain+domain] = &Host{api}

	// Initialisation du serveur pour le client
	// ----------------------------------------
	client := initClientServer()
	subdomain = viper.GetString("server.clientSubDomain")
	if subdomain != "" {
		subdomain += "."
	}
	hosts[subdomain+domain] = &Host{client}

	// Initialisation du serveur pour le Web
	// -------------------------------------
	web := initWebServer()
	subdomain = viper.GetString("server.webSubDomain")
	if subdomain != "" {
		subdomain += "."
	}
	hosts[subdomain+domain] = &Host{web}

	return hosts
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
	initApiStaticFilesAndTemplates(e)

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

	// Favicon, static files and template renderer
	// -------------------------------------------
	initClientStaticFilesAndTemplates(e)

	return e
}

// initWebServer initializes API for a client
func initWebServer() *echo.Echo {
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
	if viper.GetBool("debug.pprof") {
		protectedGroup := e.Group("")

		// Protection des routes par une Basic Auth
		// ----------------------------------------
		protectedGroup.Use(middleware.BasicAuth(
			func(username, password string, c echo.Context) (bool, error) {
				basicAuthUsername := viper.GetString("debug.basicAuthUsername")
				basicAuthPassword := viper.GetString("debug.basicAuthPassword")

				if basicAuthUsername == "" || basicAuthPassword == "" {
					return false, nil
				} else if username == basicAuthUsername && password == basicAuthPassword {
					return true, nil
				}

				return false, nil
			},
		))

		web.NewWebPprofRoute(protectedGroup).PprofRoutes()
	}

	// Liste des routes
	// ----------------
	initWebRoutes(e)

	// Favicon, static files and template renderer
	// -------------------------------------------
	initWebStaticFilesAndTemplates(e)

	return e
}
