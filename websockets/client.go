package websockets

import (
	"bytes"
	"fmt"
	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gorilla/websocket"
	"log"
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
	lib.CheckError(err, -1)
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

	// Ecoute des messages
	// -------------------
	// TODO: Faire une fonction d'aiguillage
	for {
		// Read JSON message from browser
		// ------------------------------
		var message Message
		err := c.conn.ReadJSON(&message)

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				lib.CheckError(err, 0)
			}
			break
		}

		// Print the message to the console
		fmt.Printf("%s - Message: %s with data %+v\n",
			c.conn.RemoteAddr(),
			message.Message,
			message.Data.(map[string]interface{})["text"])

		// Write message back to browser
		// -----------------------------
		err = c.conn.WriteMessage(1, []byte(message.Message))
		lib.CheckError(err, 0)

		// Broadcast example
		// -----------------
		_, messageB, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		messageB = bytes.TrimSpace(messageB)
		c.hub.broadcast <- messageB
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
