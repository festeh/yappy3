package coach

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	WS_URL  string
	URL     string
	conn    *websocket.Conn
	headers http.Header
	done    chan struct{}
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
          // TODO: broken!! fix
					// h.handleFocus(p)
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

func (h *WebSocketHandler) GetFocus() (bool, error) {
	resp, err := http.Get(h.URL)
	if err != nil {
		return false, fmt.Errorf("failed to get focus status: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Focusing bool `json:"focusing"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode focus response(%v): %v", resp.Body, err)
	}

	return result.Focusing, nil
}

func (h *WebSocketHandler) Disconnect() {
	if h.conn != nil {
		close(h.done)
		h.conn.Close()
	}
}

func (h *WebSocketHandler) FocusNow() error {
	data := "focusing=true"
	resp, err := http.Post(h.URL, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to set focus: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}