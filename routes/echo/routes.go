package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	api2 "github.com/fabienbellanger/go-rest-boilerplate/handlers/api"
	"github.com/fabienbellanger/go-rest-boilerplate/routes/api"
	"github.com/fabienbellanger/go-rest-boilerplate/routes/web"
)

// initRoutes initializes routes list
func initRoutes(e *echo.Echo) {
	// JWT configuration
	// -----------------
	jwtConfiguration := middleware.JWTConfig{
		ContextKey:  "user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		Claims:      &api2.JwtClaims{},
		SigningKey:  []byte(viper.GetString("jwt.secret")),
	}

	// Version de l'API
	// ----------------
	versionGroup := e.Group("/v1")

	// Liste des routes non protégées (à placer avant les routes protégées)
	// --------------------------------------------------------------------
	api.NewApiAuthRoute(versionGroup).AuthRoutes()
	api.NewApiBenchmarkRoute(versionGroup).BenchmarkRoutes()
	web.NewWebExampleRoute(versionGroup).ExampleRoutes()

	// Liste des routes protégées
	// --------------------------
	versionGroup.Use(middleware.JWTWithConfig(jwtConfiguration))
	api.NewApiUserRoute(versionGroup).UsersRoutes()
}
