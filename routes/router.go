package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer starts the server
func StartServer(port int) {
	// Initialisation du serveur
	// -------------------------
	router := initServer()

	// Version
	// -------
	versionGroup := router.Group("/v1");

	// Initialisation JWT
	// ------------------
	jwtMiddleware := initJWTMiddleware()

	// Liste des routes
	// ----------------
	authRoutes(versionGroup, jwtMiddleware)
	exampleRoutes(versionGroup)

	// Lancement du serveur
	// --------------------
	router.Run(":" + strconv.Itoa(port))
}

// initServer initialize the server
func initServer() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Définition de l'environnement
	// -----------------------------
	if lib.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	// ----
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"}, // TODO: Mettre dans le fichier de configuration
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Gestion des routes inconnues
	// ----------------------------
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
	})

	return router
}
