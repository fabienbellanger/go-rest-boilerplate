package gin

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// StartServer starts the server
func StartServer(port int) {
	// Initialisation du serveur
	// -------------------------
	router := initServer()

	pprof.Register(router)

	// Lancement du serveur
	// --------------------
	server := &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        router,
		ReadTimeout:    time.Duration(lib.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(lib.Config.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	go func() {
		err := server.ListenAndServe()
		if isErrorAddressAlreadyInUse(err) {
			lib.CheckError(err, 1)
		}
		lib.CheckError(err, 0)
	}()

	// Grace shutdown
	// --------------

	// Interception du signal
	// ----------------------
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	lib.GLog("Shutdown Server...")

	// Do something if you want !

	// Example
	// -------
	// Timeout (5s)
	timeout := time.Duration(lib.Config.Server.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		lib.GLog("timeout of " + timeout.String())
	}
	// End exammple

	// Shutdown the server
	// -------------------
	if err := server.Shutdown(ctx); err != nil {
		lib.CheckError(err, 0)
	}
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
	initServerEnvironment()

	// Création de l'instance de Gin
	// -----------------------------
	router := gin.New()

	if lib.Config.Log.EnableAccessLog {
		router.Use(gin.Logger())
	}
	router.Use(gin.Recovery())

	// Templates HTML
	// --------------
	router.LoadHTMLGlob("templates/**/*")

	// CORS
	// ----
	initServerCORS(router)

	// Fichiers statiques
	// ------------------
	initServerStaticFiles(router)

	// Initialisation du routing
	// -------------------------
	initServerRouting(router)

	return router
}

// initServerCORS initializes CORS
func initServerCORS(router *gin.Engine) *gin.Engine {
	router.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins:     lib.Config.Server.AllowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return router
}

// initServerStaticFiles initializes static files
func initServerStaticFiles(router *gin.Engine) *gin.Engine {
	router.Static("/js", "./assets/js")
	router.Static("/css", "./assets/css")
	router.Static("/images", "./assets/images")
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	return router
}

// initServerEnvironment initializes environment
func initServerEnvironment() {
	if lib.Config.Environment == "production" {
		// Logs dans un fichier seulement en production
		// --------------------------------------------
		gin.SetMode(gin.ReleaseMode)

		// Ouvre le fichier gin.log. S'il ne le trouve pas, il le crée
		// -----------------------------------------------------------
		logsFile, err := os.OpenFile("./"+lib.Config.Log.DirPath+lib.Config.Log.FileName, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			lib.CheckError(err, 1)
		}

		gin.DisableConsoleColor()
		gin.DefaultWriter = io.MultiWriter(logsFile)
	}
}

// initServerRouting initializes routing
func initServerRouting(router *gin.Engine) *gin.Engine {
	// Gestion des routes inconnues
	// ----------------------------
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Page not found"})
	})

	versionGroup := router.Group("/v1")
	jwtMiddleware := initJWTMiddleware()

	// Liste des routes publiques (à mettre avant les routes protégées)
	// ----------------------------------------------------------------
	exampleRoutes(versionGroup)

	// Liste des routes protégées
	// --------------------------
	authRoutes(versionGroup, jwtMiddleware)

	return router
}

// isErrorAddressAlreadyInUse checks if error is "address already in use"
func isErrorAddressAlreadyInUse(err error) bool {
	errOpError, ok := err.(*net.OpError)
	if !ok {
		return false
	}

	errSyscallError, ok := errOpError.Err.(*os.SyscallError)
	if !ok {
		return false
	}

	errErrno, ok := errSyscallError.Err.(syscall.Errno)
	if !ok {
		return false
	}

	if errErrno == syscall.EADDRINUSE {
		return true
	}

	const WSAEADDRINUSE = 10048
	if runtime.GOOS == "windows" && errErrno == WSAEADDRINUSE {
		return true
	}

	return false
}
