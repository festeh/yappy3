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
	conn       *websocket.Conn
	isFocusing bool
	headers    http.Header
}

func NewWebSocketHandler(url string) *WebSocketHandler {
	return &WebSocketHandler{
		URL:     url,
		headers: make(http.Header),
	}
}

func (h *WebSocketHandler) Connect() error {
	dialer := websocket.Dialer{}
	var err error
	
	h.conn, _, err = dialer.Dial(h.URL+"/connect", h.headers)
	if err != nil {
		return fmt.Errorf("failed to connect to websocket: %v", err)
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
