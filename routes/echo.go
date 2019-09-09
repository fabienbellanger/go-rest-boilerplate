package routes

import (
	"errors"
	"github.com/fabienbellanger/go-rest-boilerplate/handlers/user"
	"io"
	"net/http"
	"os"
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
	if lib.Config.Server.Pprof {
		pprofRoutes(e.Group(""))
	}

	// Liste des routes
	// ----------------
	initRoutes(e)

	// Favicon, static files and template renderer
	// -------------------------------------------
	initStaticFilesAndTemplates(e)

	return e
}

// initStaticFilesAndTemplates initializes static files and template renderer
func initStaticFilesAndTemplates(e *echo.Echo) {
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
}

// initCorsAndSecurity initializes CORS and Secure middlewares
func initCorsAndSecurity(e *echo.Echo) {
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
}

// initLogger initializes logger
func initLogger(e *echo.Echo) {
	lib.DefaultEchoLogWriter = os.Stdout
	if lib.Config.Environment == "production" {
		// Ouvre le fichier gin.log. S'il ne le trouve pas, il le crée
		// -----------------------------------------------------------
		logsFile, err := os.OpenFile("./"+lib.Config.Log.DirPath+lib.Config.Log.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			lib.CheckError(err, 1)
		}

		lib.DefaultEchoLogWriter = io.MultiWriter(logsFile)

		if lib.Config.Log.EnableAccessLog {
			e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
				Format:           "ECHO | ${time_custom} |  ${status} | ${latency_human}\t| ${method}\t${uri}\n",
				Output:           lib.DefaultEchoLogWriter,
				CustomTimeFormat: "2006-01-02 15:04:05",
			}))
		}
	} else {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format:           "ECHO | ${time_custom} |  ${status} | ${latency_human}\t| ${method}\t${uri}\n",
			Output:           lib.DefaultEchoLogWriter,
			CustomTimeFormat: "2006-01-02 15:04:05",
		}))
	}
}

// initRoutes initializes routes list
func initRoutes(e *echo.Echo) {
	// JWT configuration
	// -----------------
	jwtConfiguration := middleware.JWTConfig{
		ContextKey:  "user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		Claims:      &user.JwtClaims{},
		SigningKey:  []byte(lib.Config.Jwt.Secret),
	}

	// Version de l'API
	// ----------------
	versionGroup := e.Group("/v1")

	// Liste des routes non protégées (à placer avant les routes protégées)
	// --------------------------------------------------------------------
	authRoutes(versionGroup)
	exampleRoutes(versionGroup)

	// Liste des routes protégées
	// --------------------------
	versionGroup.Use(middleware.JWTWithConfig(jwtConfiguration))
	usersRoutes(versionGroup)
}

// customHTTPErrorHandler manages HTTP errors
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
	}

	switch code {
	case http.StatusUnauthorized:
		// 401
		c.JSON(code, map[string]string{"message": "Not Authorized"})
	case http.StatusNotFound:
		// 404
		c.JSON(code, map[string]string{"message": "Resource Not Found"})
	case http.StatusInternalServerError:
		// 500
		c.JSON(code, map[string]string{"message": "Internal Server Error"})
	default:
		c.JSON(code, "")
	}
}
