package routes

import (
	"errors"
	"io"
	"net/http"
	"os"
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
	versionGroup := router.Group("/v1")

	// Initialisation JWT
	// ------------------
	jwtMiddleware := initJWTMiddleware()

	// Liste des routes publiques (à mettre avant les routes protégées)
	// ----------------------------------------------------------------
	exampleRoutes(versionGroup)

	// Liste des routes protégées
	// --------------------------
	authRoutes(versionGroup, jwtMiddleware)

	// Lancement du serveur
	// --------------------
	err := router.Run(":" + strconv.Itoa(port))
	lib.CheckError(err, -1)
}

// initServer initialize the server
func initServer() *gin.Engine {
	// La config pour le CORS est-elle correcte ?
	// ------------------------------------------
	if !lib.IsServerConfigCorrect() {
		lib.CheckError(errors.New("no allow origins defined in config file"), 1)
	}

	// Définition de l'environnement
	// -----------------------------
	if lib.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)

		// Logs dans un fichier seulement en production
		// --------------------------------------------

		// Ouvre le fichier gin.log. S'il ne le trouve pas, il le crée
		// -----------------------------------------------------------

		logsFile, err := os.OpenFile("./"+lib.Config.Log.DirPath+lib.Config.Log.FileName, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			lib.CheckError(err, -1)
		}

		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(logsFile)
	}

	// Création de l'instance de Gin
	// -----------------------------
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS
	// ----
	router.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins:     lib.Config.Server.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Gestion des routes inconnues
	// ----------------------------
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
	})

	// Templates HTML
	// --------------
	router.LoadHTMLGlob("templates/**/*")

	// Fichiers statiques
	// ------------------
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	return router
}
