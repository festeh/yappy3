package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for development
	},
}

type WebSocketHandler struct {
	URL string
}

func NewWebSocketHandler(url string) *WebSocketHandler {
	return &WebSocketHandler{
		URL: url,
	}
}

func (h *WebSocketHandler) Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		// Handle different message types
		switch string(p) {
		case "focus":
			h.handleFocus(conn)
		default:
			log.Printf("Unknown message: %s", string(p))
		}

		// Echo the message back
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error writing message: %v", err)
			return
		}
	}
}

func (h *WebSocketHandler) handleFocus(conn *websocket.Conn) {
	// Dummy implementation for focus event
	log.Println("Focus event received")
	msg := fmt.Sprintf("Focused on %s", h.URL)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.Printf("Error sending focus response: %v", err)
	}
}
