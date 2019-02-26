package routes

import (
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-gonic/gin"
)

// StartServer : Démarrage du serveur
func StartServer(port int) {
	// Initialisation du serveur
	// e := initServer()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if lib.Config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

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
// func initServer() *echo.Echo {
// 	e := echo.New()

// 	// Logger
// 	logFile, err := os.OpenFile("api-access.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
// 	defer logFile.Close()

// 	if err != nil {
// 		log.Errorf("Cannot open 'api-access.log', (%s)", err.Error())
// 		flag.Usage()
// 		os.Exit(-1)
// 	}

// 	e.Debug = true

// 	// if e.Debug {
// 	// 	e.Logger.SetLevel(log.INFO)
// 	// }

// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 		Format: "${time_rfc3339} |  ${status} | ${latency_human}\t| ${method}\t${uri}\n",
// 		Output: logFile,
// 	}))

// 	// Recover
// 	e.Use(middleware.Recover())

// 	// CORS
// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
// 	}))

// 	// Secure
// 	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
// 		XSSProtection:         "1; mode=block",
// 		ContentTypeNosniff:    "nosniff",
// 		XFrameOptions:         "SAMEORIGIN",
// 		HSTSMaxAge:            3600,
// 		ContentSecurityPolicy: "default-src 'self'",
// 	}))

// 	// Favicon
// 	e.File("/favicon.ico", "assets/images/favicon.ico")

// 	return e
// }
