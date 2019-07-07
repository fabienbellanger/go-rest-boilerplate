package routes

import (
	"net/http"
	"time"

	jwtEcho "github.com/dgrijalva/jwt-go"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/labstack/echo/v4"
)

// echoAuthRoutes manages authentication routes for Echo
func echoAuthRoutes(e *echo.Echo, g *echo.Group) {
	// Liste des routes
	g.POST("/login", loginHandler)
}

func loginHandler(c echo.Context) error {
	// Récupération des variables transmises
	// -------------------------------------
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Vérification en base
	// --------------------
	user, err := models.CheckLogin(username, password)

	if err == nil && user.ID != 0 {
		// Création du token d'authentification
		// ------------------------------------
		token := jwtEcho.New(jwtEcho.SigningMethodHS256)

		// Enregistrement de la revendication
		// ----------------------------------
		claims := token.Claims.(jwtEcho.MapClaims)
		claims["id"] = user.ID
		claims["lastname"] = user.Lastname
		claims["firstname"] = user.Firstname
		claims["exp"] = time.Now().Add(time.Hour).Unix()

		// Génération du token encodé et envoi dans la réponse
		// ---------------------------------------------------
		t, err := token.SignedString([]byte(lib.Config.Jwt.Secret))

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":     t,
			"id":        user.ID,
			"lastname":  user.Lastname,
			"firstname": user.Firstname,
			"fullname":  user.GetFullname(),
		})
	}

	return echo.ErrUnauthorized
}
