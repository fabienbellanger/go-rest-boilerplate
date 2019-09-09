package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/fabienbellanger/go-rest-boilerplate/handlers/user"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/routes/api"
)

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
	api.NewApiAuthRoute(versionGroup).AuthRoutes()
	api.NewApiExampleRoute(versionGroup).ExampleRoutes()

	// Liste des routes protégées
	// --------------------------
	versionGroup.Use(middleware.JWTWithConfig(jwtConfiguration))
	api.NewApiUserRoute(versionGroup).UsersRoutes()
}
