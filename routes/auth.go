package routes

import (
	"github.com/appleboy/gin-jwt"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

// authRoutes : Partie authentification
func authRoutes(group *gin.RouterGroup, jwtMiddleware *jwt.GinJWTMiddleware) {
	// Liste des routes
	// ----------------
	group.POST("/login", jwtMiddleware.LoginHandler)
	group.GET("/refresh-token", jwtMiddleware.RefreshHandler)

	group.Use(jwtMiddleware.MiddlewareFunc())
	{
		group.GET("/hello", helloHandler)
	}
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	c.JSON(http.StatusOK, lib.GetHTTPResponse(
		http.StatusOK,
		"Success",
		gin.H{
			"id":        claims["id"],
			"username":  claims["username"],
			"lastname":  claims["lastname"],
			"firstname": claims["firstname"],
			"fullname":  claims["fullname"],
			"text":      "Hello World.",
		}),
	)
}
