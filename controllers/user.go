package controllers

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/labstack/echo/v4"
)

// JwtClaims are custom claims extending default ones.
type JwtClaims struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	jwt.StandardClaims
}

// userLogin is used for binding data in login route
type userLogin struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

// GetUserDetailsHandler displays authenticated user information
func GetUserDetailsHandler(c echo.Context) error {
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid parameters",
		})
	}

	// Vérification en base
	// --------------------
	user, err := models.CheckLogin(u.Username, u.Password)
	if err != nil || user.ID == 0 {
		return echo.ErrUnauthorized
	}

	// Enregistrement de la revendication
	// ----------------------------------
	claims := &JwtClaims{
		user.ID,
		user.Username,
		user.Password,
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

// ChangePassword changes user password
func ChangePassword(c echo.Context) error {
	type data struct {
		CurrentPassword    string `json:"currentPassword" form:"currentPassword" query:"currentPassword"`
		NewPassword        string `json:"newPassword" form:"newPassword" query:"newPassword"`
		ConfirmNewPassword string `json:"confirmNewPassword" form:"confirmNewPassword" query:"confirmNewPassword"`
	}

	// Récupération des variables transmises
	// -------------------------------------
	input := new(data)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid parameters",
		})
	}

	// Récupération du claims
	// ----------------------
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	newPassword := fmt.Sprintf("%x", sha512.Sum512([]byte(input.NewPassword)))

	// Tests validité des paramètres
	// -----------------------------
	if claims.Password == newPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "New password must be different from old password",
		})
	}

	if len(input.NewPassword) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "New password must contain at least 8 caracters",
		})
	}

	if input.NewPassword != input.ConfirmNewPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Confirm password must be the same as new password",
		})
	}

	return errors.New("error")
}
