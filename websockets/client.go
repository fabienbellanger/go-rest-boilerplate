package websockets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

// Connection parameters
var ws = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Duration(time.Second * 5),
}

// Message represents the general message structure
type Message struct {
	Message string
	Data    interface{}
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection
	conn *websocket.Conn

	// Buffered channel of outbound messages
	sendBroadcast chan []byte

	// Id of the client (type of client: account, terminal, etc.)
	id string
}

// ClientConnection connects client to the server
func ClientConnection(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// Connexion
	// ---------
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		lib.CheckError(err, 0)
		return
	}

	fmt.Println("Connexion du client...")

	// Création du client
	// ------------------
	client := &Client{hub: hub, conn: conn, sendBroadcast: make(chan []byte, 256), id: "message"}
	client.hub.register <- client

	// Ecoute des messages
	// -------------------
	go client.readMessages()

	// Broadcast des messages sur le hub
	// ---------------------------------
	go client.broadcastMessage()
}

// readMessages reads message for server
func (c *Client) readMessages() {
	defer func() {
		// Déconnexion du hub
		// ------------------
		c.hub.unregister <- c

		// Déconnexion du hub
		// ------------------
		err := c.conn.Close()
		lib.CheckError(err, 0)
	}()

	// Gestion des messages
	// --------------------
	c.manageMessages()
}

// manageMessages manages sending and writing messages
func (c *Client) manageMessages() {
	for {
		// Read message from browser
		// -------------------------
		_, messageStr, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// Déconnexion du client
				// ---------------------
				lib.CheckError(err, 0)
			}

			break
		}

		// Trim de la chaîne
		// -----------------
		messageStr = bytes.TrimSpace(messageStr)

		// Est-ce un JSON ?
		// ----------------
		var messageJSON Message
		err = json.Unmarshal(messageStr, &messageJSON)

		if err != nil {
			// Not a JSON message
			// ------------------

			// Par exemple, on broadcast le message
			messageStr = bytes.TrimSpace(messageStr)
			c.hub.broadcast <- messageStr

		} else {
			// JSON message
			// ------------

			// Par exemple
			// Print the message to the console
			fmt.Printf("%s - Message: %s with data %+v\n",
				c.conn.RemoteAddr(),
				messageJSON.Message,
				messageJSON.Data.(map[string]interface{})["text"])

			// Write message back to browser
			// -----------------------------
			err = c.conn.WriteMessage(1, []byte(messageJSON.Message))
			lib.CheckError(err, 0)
		}
	}
}

// broadcastMessage writes message to all clients of the hub
func (c *Client) broadcastMessage() {
	defer func() {
		// Déconnexion du hub
		// ------------------
		err := c.conn.Close()
		lib.CheckError(err, 0)
	}()

	// Envoi des messages
	// ------------------
	for {
		select {
		case message, ok := <-c.sendBroadcast:
			if !ok {
				// The hub closed the channel
				err := c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				lib.CheckError(err, 0)

				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			_, err = w.Write(message)
			lib.CheckError(err, 0)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
