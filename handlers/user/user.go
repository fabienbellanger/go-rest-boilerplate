package user

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/models"
	"github.com/fabienbellanger/go-rest-boilerplate/repositories"
	userRepository "github.com/fabienbellanger/go-rest-boilerplate/repositories/user"
)

// JwtClaims are custom claims extending default ones.
type JwtClaims struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	jwt.StandardClaims
}

// Handler
type Handler struct {
	repository repositories.UserRepository
}

// NewUserHandler
func NewUserHandler() *Handler {
	return &Handler{
		repository: userRepository.NewMysqlUserRepository(),
	}
}

// GetUserDetailsHandler displays authenticated user information
func (h *Handler) GetUserDetailsHandler(c echo.Context) error {
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
func (h *Handler) LoginHandler(c echo.Context) error {
	type userLogin struct {
		Username string `json:"username" form:"username" query:"username"`
		Password string `json:"password" form:"password" query:"password"`
	}

	// Récupération des variables transmises
	// -------------------------------------
	input := new(userLogin)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid parameters",
		})
	}

	// Vérification en base
	// --------------------
	user, err := h.repository.CheckLogin(input.Username, input.Password)
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
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(viper.GetInt("jwt.expirationTime"))).Unix(),
		},
	}

	// Création du token d'authentification
	// ------------------------------------
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Génération du token encodé et envoi dans la réponse
	// ---------------------------------------------------
	t, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
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
func (h *Handler) ChangePassword(c echo.Context) error {
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
	u := c.Get("user").(*jwt.Token)
	claims := u.Claims.(*JwtClaims)
	newPassword := fmt.Sprintf("%x", sha512.Sum512([]byte(input.NewPassword)))

	// Récupération de l'utilisateur en base
	// -------------------------------------
	var user models.User
	database.Orm.First(&user, claims.ID)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}

	// Tests validité des paramètres
	// -----------------------------
	if user.Password == newPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "New password must be different from old password",
		})
	} else if len(input.NewPassword) < 8 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "New password must contain at least 8 caracters",
		})
	} else if input.NewPassword != input.ConfirmNewPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Confirm password must be the same as new password",
		})
	}

	// Modification en base
	// --------------------
	ok := h.repository.ChangePassword(&user, newPassword)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "An error has occured",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Password changes with success",
	})
}
