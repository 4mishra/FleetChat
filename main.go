package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

// Upgrader for WebSocket connections
var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	clients = make(map[*websocket.Conn]bool) // Store active connections
	broadcast = make(chan []byte) // Channel for broadcasting messages
)


// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade connection to WebSocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Add the new connection to the clients map
	clients[conn] = true

	// Handle WebSocket messages
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			// Remove the connection from the clients map
			delete(clients, conn)
			return
		}

		// Process the received message based on messageType
		// Example: Handle text messages
		switch messageType {
		case websocket.TextMessage:
			// text := string(p)
			// Echo the received message
			err = conn.WriteMessage(messageType, p)
			if err != nil {
				return
			}
		case websocket.BinaryMessage:
		// Handle binary message (p is a []byte)
		// Perform actions based on the received binary message
		// ...

		case websocket.CloseMessage:
			// Handle close message
			// ...
		}
		if messageType == websocket.TextMessage {
			// Process the message (p)
			// ...
			// Echo the message back to the client
			err = conn.WriteMessage(messageType, p)
			if err != nil {
				return
			}
		}
	}
}


func BroadcastMessages() {
	for {
		// Receive message from the broadcast channel
		message := <-broadcast

		// Broadcast the message to all active connections
		for client := range clients {
			if err := client.WriteMessage(messageType, p); err != nil {
				// If sending message to client fails, remove the connection
				delete(clients, client)
				client.Close()
			}
		}
	}
}
func main() {
	// Start the broadcast goroutine
	go BroadcastMessages()
	http.HandleFunc("/ws", HandleWebSocket)
	http.ListenAndServe(":8080", nil)
}
