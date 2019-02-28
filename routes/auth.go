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
}
