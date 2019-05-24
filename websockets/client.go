package websockets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fabienbellanger/go-rest-boilerplate/lib"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
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

	sendMessage chan Message

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
	client := &Client{hub: hub, conn: conn, sendBroadcast: make(chan []byte, 256), sendMessage: make(chan Message), id: "message"}
	client.hub.register <- client

	// Envoi des messages
	// ------------------
	go client.writeMessage()

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
			// Not a valid JSON message
			// ------------------------
			lib.CheckError(err, 0)
		} else {
			// Correct JSON message
			// --------------------
			if messageJSON.Message != "" {
				c.sendMessage <- messageJSON
			}
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
				lib.CheckError(err, 0)

				return
			}

			_, err = w.Write(message)
			lib.CheckError(err, 0)

			if err := w.Close(); err != nil {
				lib.CheckError(err, 0)

				return
			}
		}
	}
}

// writeMessage writes message to socket connection
func (c *Client) writeMessage() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message := <-c.sendMessage:
			switch message.Message {
			case "test":
				c.test(message)
			}
		}
	}
}

// test is a test of websockets
func (c *Client) test(message Message) {
	// Write message back to browser
	// -----------------------------
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message.Message))
	lib.CheckError(err, 0)

	type testType struct {
		Text struct {
			Toto string
		}
	}

	var t testType
	err = mapstructure.Decode(message.Data, &t)
	if err != nil {
		lib.CheckError(err, 0)
	}

	fmt.Printf("%#v - %s\n", t, t.Text.Toto)

	// Broadcast du message (attention problème de concurrence, ne pas faire plusieurs write en même temps)
	// ----------------------------------------------------------------------------------------------------
	// c.hub.broadcast <- []byte(message.Message)
}
