package routes

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/fabienbellanger/go-rest-boilerplate/modules/user"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// loginParametersType represents login parameters
type loginParametersType struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// initJWTMiddleware initialize JWT middleware
func initJWTMiddleware() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte(lib.Config.Jwt.Secret),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user.UserDB); ok {
				return jwt.MapClaims{
					"id":        v.ID,
					"username":  v.Username,
					"lastname":  v.Lastname,
					"firstname": v.Firstname,
					"fullname":  v.GetFullname(),
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
			userToCheck, err := user.CheckLogin(username, password)

			if err == nil && userToCheck.ID != 0 {
				return &userToCheck, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// data semble correspondre au claims["id"]
			// fmt.Printf("Authorizator %v of type %v\n", data, reflect.TypeOf(data))

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
