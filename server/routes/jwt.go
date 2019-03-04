package routes

import (
	"apiticSellers/server/database"
	"apiticSellers/server/lib"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// loginParametersType represents login parameters
type loginParametersType struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserDB type
type UserDB struct {
	ID        uint64
	Username  string
	Password  string
	Lastname  string
	Firstname string
	CreatedAt sql.RawBytes
	DeletedAt sql.RawBytes
}

// initJWTMiddleware initialize JWT middleware
func initJWTMiddleware() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte(lib.Config.Jwt.Secret),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// PayloadFunc maps the claims in the JWT
			// --------------------------------------
			if v, ok := data.(*UserDB); ok {
				return jwt.MapClaims{
					"id":        v.ID,
					"username":  v.Username,
					"lastname":  v.Lastname,
					"firstname": v.Firstname,
					"fullname":  v.getFullname(),
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginParameters loginParametersType

			if err := c.ShouldBind(&loginParameters); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			username := loginParameters.Username
			password := loginParameters.Password

			// Vérification en base
			// --------------------
			userToCheck, err := checkLogin(username, password)

			if err == nil && userToCheck.ID != 0 {
				return &userToCheck, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// INFO: data semble correspondre au claims["id"]

			// Si un ID existe, l'utilisateur est autorisé à utiliser l'API
			if v := uint64(data.(float64)); v > 0 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, lib.GetHTTPResponse(
				code,
				"Unauthorized: "+message,
				nil,
			))
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	return
}

// checkLogin checks if username and password are corrct
func checkLogin(username, password string) (UserDB, error) {
	encryptPassword := sha512.Sum512([]byte(password))
	encryptPasswordStr := hex.EncodeToString(encryptPassword[:])
	query := `
		SELECT id, username, lastname, firstname, created_at, deleted_at
		FROM user
		WHERE username = ? AND password = ? AND deleted_at IS NULL
		LIMIT 1`
	rows, err := database.Select(query, username, encryptPasswordStr)

	var user UserDB

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Lastname, &user.Firstname, &user.CreatedAt, &user.DeletedAt)

		lib.CheckError(err, 0)
	}

	return user, err
}

// getFullname displays user fullname
func (u UserDB) getFullname() string {
	return u.Firstname + " " + u.Lastname
}
