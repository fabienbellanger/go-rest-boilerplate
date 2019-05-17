package routes

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var ws = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

// WebsocketsServerStart starts websockets server
func WebsocketsServerStart() {
	http.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := ws.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":8082", nil)
}
