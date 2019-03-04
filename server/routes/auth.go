package routes

import (
	"apiticSellers/server/lib"
	"github.com/appleboy/gin-jwt"
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
		group.GET("/users", getUserHandler)
	}
}

// getUserHandler displays authenticated user information
// TODO: Mettre dans un controleur
func getUserHandler(c *gin.Context) {
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
		}),
	)
}
