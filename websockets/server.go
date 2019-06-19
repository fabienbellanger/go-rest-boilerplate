package websockets

import (
	"net/http"
	"strconv"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
)

// ServerStart starts websockets server
func ServerStart(port int) {
	// Lancement du hub
	// ----------------
	hub := newHub()
	go hub.run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ClientConnection(hub, w, r)
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	lib.CheckError(err, 1)

	// Pour utiliser les wss (WebSocket Secure)
	// ----------------------------------------
	// err := http.ListenAndServeTLS(":8443", lib.Config.SSL.CertPath, lib.Config.SSL.KeyPath, nil)
	// lib.CheckError(err, 1)
}
