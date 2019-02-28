package routes

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
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

	c.JSON(200, gin.H{
		"userID": claims["id"],
		"user":   claims["user"],
		"text":   "Hello World.",
	})
}
