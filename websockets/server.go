package websockets

import (
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ServerStart starts websockets server
func ServerStart(port int) {
	// Lancement du hub
	// ----------------
	hub := newHub()
	go hub.run()

	e := echo.New()

	// Logger
	// ------
	if lib.Config.Environment == "development" {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[WS] ${time_rfc3339} |  ${status} | ${latency_human}\t| ${method}\t${uri}\n",
		}))
	}

	// Recover
	// -------
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/", func(c echo.Context) error {
		ClientConnection(hub, c.Response(), c.Request())

		return nil
	})
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	ClientConnection(hub, w, r)
	// })

	// err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	// lib.CheckError(err, 1)

	// Pour utiliser les wss (WebSocket Secure)
	// ----------------------------------------
	// err := http.ListenAndServeTLS(":8443", lib.Config.SSL.CertPath, lib.Config.SSL.KeyPath, nil)
	// lib.CheckError(err, 1)
}
