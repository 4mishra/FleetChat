package main

import (
    "net/http"
    "github.com/gorilla/websocket"
)

// Upgrader for WebSocket connections
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// HandleWebSocket handles WebSocket connections
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    // Upgrade HTTP connection to WebSocket connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        http.Error(w, "Could not upgrade connection to WebSocket", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    // Handle WebSocket messages
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }

        // Process the received message based on messageType
        // Example: Handle text messages
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

func main() {
    http.HandleFunc("/ws", HandleWebSocket)
    http.ListenAndServe(":8080", nil)
}
