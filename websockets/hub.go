package websockets

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// newHub initializes new Hub
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// run runs the hub
func (h *Hub) run() {
	for {
		select {
		// Connexion
		// ---------
		case client := <-h.register:
			h.clients[client] = true

		// Déconnexion
		// -----------
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				// Suppression du client de la liste
				delete(h.clients, client)

				// Fermeture du chan
				close(client.sendBroadcast)
			}

		// Broadcast du message à tous les clients
		// ---------------------------------------
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.sendBroadcast <- message:
				default:
					// Fermeture du chan
					close(client.sendBroadcast)

					// Suppression du client de la liste
					delete(h.clients, client)
				}
			}
		}
	}
}
