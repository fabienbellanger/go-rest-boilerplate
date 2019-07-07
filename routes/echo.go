package routes

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// StartEchoServer starts Echo web server
func StartEchoServer(port int) {
	e := initEchoServer()

	// Start server
	// ------------
	s := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  time.Duration(lib.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(lib.Config.Server.WriteTimeout) * time.Second,
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
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "SAMEORIGIN",
		HSTSMaxAge:         3600,
		// ContentSecurityPolicy: "default-src 'self'",
	}))

	// Version de l'API
	// ----------------
	versionGroup := e.Group("/v1")

	// JWT configuration
	// -----------------
	jwtConfiguration := middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: []byte(lib.Config.Jwt.Secret),
	}

	// Profilage
	// ---------
	if lib.Config.Server.Pprof {
		pprofRoutes(versionGroup)
	}

	// Liste des routes non protégées
	// ------------------------------
	authRoutes(e, versionGroup)
	exampleRoutes(e, versionGroup)

	// Liste des routes protégées
	// --------------------------
	versionGroup.Use(middleware.JWTWithConfig(jwtConfiguration))

	versionGroup.GET("/restricted", func(c echo.Context) error {
		return c.String(http.StatusOK, "Restricted page")
	})

	// Favicon
	// -------
	e.File("/favicon.ico", "assets/favicon.ico")

	// Assets
	// ------
	e.Static("/js", "./assets/js")
	e.Static("/css", "./assets/css")
	e.Static("/images", "./assets/images")

	// Templates
	// ---------
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/**/*.gohtml")),
	}
	e.Renderer = t

	return e
}

func initJWT() {

}
