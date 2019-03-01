package routes

import (
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// loginParametersType represents login parameters
type loginParametersType struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User structure
type User struct {
	ID        uint64
	UserName  string
	FirstName string
	LastName  string
}

// initJWTMiddleware initialize JWT middleware
func initJWTMiddleware() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware = &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte(lib.Config.Jwt.Secret), // TODO: Récupérer depuis le fichier de config
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					"id":   v.ID,
					"user": v,
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

			// TODO: Test d'exemple, à modifier
			if username == "admin" && password == "pwd" {
				return &User{
					ID:        123,
					UserName:  username,
					LastName:  "Admin",
					FirstName: "Admin",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// data semble correspondre au claims["id"]
			// fmt.Printf("Authorizator %v of type %v\n", data, reflect.TypeOf(data))

			// TODO: Test d'exemple, à modifier
			if v := uint64(data.(float64)); v == 123 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, lib.GetHTTPResponse(
				code,
				"Unauthorized: " + message,
				nil,
			))
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	return
}
