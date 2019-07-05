package routes

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	jwtEcho "github.com/dgrijalva/jwt-go"
	"github.com/fabienbellanger/go-rest-boilerplate/controllers"
	"github.com/fabienbellanger/go-rest-boilerplate/database"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/orm/models"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

// checkLogin checks if username and password are correct
// TODO: Utiliser GORM
func checkLogin(username, password string) (models.User, error) {
	encryptPassword := sha512.Sum512([]byte(password))
	encryptPasswordStr := hex.EncodeToString(encryptPassword[:])
	query := `
		SELECT id, username, lastname, firstname, created_at, deleted_at
		FROM users
		WHERE username = ? AND password = ? AND deleted_at IS NULL
		LIMIT 1`
	rows, err := database.Select(query, username, encryptPasswordStr)

	var user models.User

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Lastname, &user.Firstname, &user.CreatedAt, &user.DeletedAt)

		lib.CheckError(err, 0)
	}

	return user, err
}

// authRoutes manages authentication routes
func authRoutes(group *gin.RouterGroup, jwtMiddleware *jwt.GinJWTMiddleware) {
	// Liste des routes
	// ----------------
	group.POST("/login", jwtMiddleware.LoginHandler)
	group.GET("/refresh-token", jwtMiddleware.RefreshHandler)

	group.Use(jwtMiddleware.MiddlewareFunc())
	{
		group.GET("/users", controllers.GetUserHandler)
	}
}

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
	user, err := checkLogin(username, password)

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
