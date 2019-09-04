package controllers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/labstack/echo/v4"
)

// JwtClaims are custom claims extending default ones.
type JwtClaims struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	jwt.StandardClaims
}

// userLogin is used for binding data in login route
type userLogin struct {
	username string `json:"username" form:"username" query:"username"`
	password string `json:"password" form:"password" query:"password"`
}

// GetUserHandler displays authenticated user information
func GetUserHandler(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id":        claims.ID,
		"username":  claims.Username,
		"lastname":  claims.Lastname,
		"firstname": claims.Firstname,
	})
}

// LoginHandler make authentication
func LoginHandler(c echo.Context) error {
	// Récupération des variables transmises
	// -------------------------------------
	u := new(userLogin)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Vérification en base
	// --------------------
	user, err := models.CheckLogin(u.username, u.password)
	if err != nil || user.ID == 0 {
		return echo.ErrUnauthorized
	}

	// Enregistrement de la revendication
	// ----------------------------------
	claims := &JwtClaims{
		user.ID,
		user.Username,
		user.Lastname,
		user.Firstname,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lib.Config.Jwt.ExpirationTime)).Unix(),
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
		"expireAt":  time.Unix(claims.StandardClaims.ExpiresAt, 0),
	})
}
