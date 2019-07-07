package routes

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/labstack/echo/v4"
)

// JwtClaims are custom claims extending default ones.
type JwtClaims struct {
	ID        uint64 `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	jwt.StandardClaims
}

// authRoutes manages authentication routes for Echo
func authRoutes(e *echo.Echo, g *echo.Group) {
	// Liste des routes
	g.POST("/login", loginHandler)
}

// loginHandler make authentication
func loginHandler(c echo.Context) error {
	// Récupération des variables transmises
	// -------------------------------------
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Vérification en base
	// --------------------
	user, err := models.CheckLogin(username, password)

	if err != nil || user.ID == 0 {
		return echo.ErrUnauthorized
	}

	// Enregistrement de la revendication
	// ----------------------------------
	claims := &JwtClaims{
		user.ID,
		user.Lastname,
		user.Firstname,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Création du token d'authentification
	// ------------------------------------
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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
