package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	WS_URL     string
	URL        string
	conn       *websocket.Conn
	headers    http.Header
	focusing   bool
	OnFocusSet func(bool)
	done       chan struct{}
}

func (h *WebSocketHandler) setFocusing(focusing bool) {
	h.focusing = focusing
	if h.OnFocusSet != nil {
		h.OnFocusSet(focusing)
	}
}

func NewWebSocketHandler(wsurl string, url string) *WebSocketHandler {
	return &WebSocketHandler{
		WS_URL:  wsurl,
		URL:     url,
		headers: make(http.Header),
		done:    make(chan struct{}),
	}
}

func (h *WebSocketHandler) Connect() error {
	dialer := websocket.Dialer{}
	var err error

	h.conn, _, err = dialer.Dial(h.WS_URL, h.headers)
	if err != nil {
		log.Println("ws_url", h.WS_URL)
		return fmt.Errorf("failed to connect to websocket: %v", err)
	}

	go func() {
		defer h.conn.Close()

		for {
			select {
			case <-h.done:
				return
			default:
				messageType, p, err := h.conn.ReadMessage()
				if err != nil {
					log.Printf("Error reading message: %v", err)
					return
				}

				// Parse JSON message
				var msg struct {
					Event string `json:"event"`
				}
				if err := json.Unmarshal(p, &msg); err != nil {
					log.Printf("Error parsing message: %v", err)
					continue
				}

				// Handle different event types
				switch msg.Event {
				case "focusing":
					h.handleFocus(p)
				default:
					log.Printf("Unknown event: %s", msg.Event)
				}

				// Echo the message back
				if err := h.conn.WriteMessage(messageType, p); err != nil {
					log.Printf("Error writing message: %v", err)
					return
				}
			}
		}
	}()

	return nil
}

func (h *WebSocketHandler) GetFocus() error {
	resp, err := http.Get(h.URL)
	if err != nil {
		return fmt.Errorf("failed to get focus status: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Focusing bool `json:"focusing"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode focus response(%v): %v", resp.Body, err)
	}

	h.setFocusing(result.Focusing)
	return nil
}

func (h *WebSocketHandler) Disconnect() {
	if h.conn != nil {
		close(h.done)
		h.conn.Close()
	}
}

func (h *WebSocketHandler) handleFocus(message []byte) {
	// Parse the message to get focus state
	var msg struct {
		Event string `json:"event"`
		Focusing bool   `json:"focusing"`
	}
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Printf("Error parsing focus message: %v", err)
		return
	}

	h.setFocusing(msg.Focusing)
	log.Printf("Focus state updated: %v", h.focusing)
}
