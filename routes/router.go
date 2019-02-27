package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// StartServer : Démarrage du serveur
func StartServer(port int) {
	// Initialisation du serveur
	router := initServer()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	// :param : Paramètre obligatoire
	// *param : Paramètre optionnel
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":" + strconv.Itoa(port))
}

// initServer : Initialisation du serveur
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
