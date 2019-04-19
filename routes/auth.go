package routes

import (
	"github.com/appleboy/gin-jwt"
	"github.com/fabienbellanger/go-rest-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

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
