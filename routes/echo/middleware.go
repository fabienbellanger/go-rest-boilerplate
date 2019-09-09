package echo

import (
	"io"
	"os"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
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
