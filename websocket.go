package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	URL        string
	upgrader   websocket.Upgrader
	conn       *websocket.Conn
	isFocusing bool
}

func NewWebSocketHandler(url string) *WebSocketHandler {
	return &WebSocketHandler{
		URL: url,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins for development
			},
		},
	}
}

func (h *WebSocketHandler) Connect() {
	var err error
	h.conn, err = h.upgrader.Upgrade(nil, nil, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	for {
		messageType, p, err := h.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}

		// Handle different message types
		switch string(p) {
		case "focus":
			h.handleFocus()
		default:
			log.Printf("Unknown message: %s", string(p))
		}

		// Echo the message back
		if err := h.conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error writing message: %v", err)
			return
		}
	}
}

func (h *WebSocketHandler) GetFocus() error {
	resp, err := http.Get(h.URL + "/focus")
	if err != nil {
		return fmt.Errorf("failed to get focus status: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Focusing bool `json:"focusing"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode focus response: %v", err)
	}

	h.isFocusing = result.Focusing
	return nil
}

func (h *WebSocketHandler) handleFocus() {
	// Dummy implementation for focus event
	log.Println("Focus event received")
	msg := fmt.Sprintf("Focused on %s", h.URL)
	if err := h.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		log.Printf("Error sending focus response: %v", err)
	}
}
