package routes

import (
	"fmt"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"time"
)

var ws = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

// WebSocketServerStart starts websockets server
func WebSocketServerStart(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := ws.Upgrade(w, r, nil)
		lib.CheckError(err, -1)

		fmt.Println("Connexion au client...")

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				// La connexion est perdue
				// -----------------------
				lib.CheckError(err, 0)

				break
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			err = conn.WriteMessage(msgType, msg)
			lib.CheckError(err, 0)
		}
	})

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	lib.CheckError(err, -1)
}
