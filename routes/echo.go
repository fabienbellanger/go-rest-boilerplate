package routes

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartEchoServer(port int) {
	// La config pour le CORS est-elle correcte ?
	// ------------------------------------------
	if !lib.IsServerConfigCorrect() {
		lib.CheckError(errors.New("no allow origins defined in config file"), 1)
	}

	// Echo instance
	e := echo.New()

	// Logger
	// ------
	if lib.Config.Environment == "development" {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[ECHO] ${time_rfc3339} |  ${status} | ${latency_human}\t| ${method}\t${uri}\n",
		}))
	}

	// Recover
	// -------
	e.Use(middleware.Recover())

	// CORS
	// ----
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     lib.Config.Server.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))

	// Secure
	// ------
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	// Version de l'API
	// ----------------
	versionGroup := e.Group("/v1/echo")

	// Liste des routes
	echoAuthRoutes(e, versionGroup)
	echoExampleRoutes(e, versionGroup)

	// Favicon
	// -------
	e.Static("/js", "./assets/js")
	e.Static("/css", "./assets/css")
	e.Static("/images", "./assets/images")
	e.File("/favicon.ico", "assets/images/favicon.ico")

	// Start server
	// ------------
	s := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
	// e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
