package websockets

import (
	"fmt"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
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

type Message struct {
	Message string
	Data    interface{}
}

// ClientConnection connects client to the server
func ClientConnection(w http.ResponseWriter, r *http.Request) {
	// Connexion
	// ---------
	conn, err := ws.Upgrade(w, r, nil)
	lib.CheckError(err, -1)
	fmt.Println("Connexion au client...")

	for {
		// Read JSON message from browser
		// ------------------------------
		var message Message
		err = conn.ReadJSON(&message)

		if err != nil {
			lib.CheckError(err, 0)

			break
		}

		// Print the message to the console
		fmt.Printf("2 -> %s - Message: %s with data %+v\n", conn.RemoteAddr(), message.Message, message.Data.(map[string]interface{})["text"])

		// Write message back to browser
		// -----------------------------
		err = conn.WriteMessage(1, []byte(message.Message))
		lib.CheckError(err, 0)
	}
}
