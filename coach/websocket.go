package coach

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
)

import (
	"sync"
)

type WebSocketHandler struct {
	WS_URL    string
	URL       string
	conn      *websocket.Conn
	headers   http.Header
	done      chan struct{}
	msgChan   chan []byte
	mu        sync.Mutex
	connected bool
}

func NewWebSocketHandler(wsurl string, url string) *WebSocketHandler {
	return &WebSocketHandler{
		WS_URL:  wsurl,
		URL:     url,
		headers: make(http.Header),
		done:    make(chan struct{}),
		msgChan: make(chan []byte, 5),
	}
}

func (h *WebSocketHandler) Connect() error {
	h.mu.Lock()
	if h.connected {
		h.mu.Unlock()
		return nil
	}

	var err error
	err = h.connectWebSocket()
	if err != nil {
		h.mu.Unlock()
		return fmt.Errorf("failed to connect to websocket: %v", err)
	}

	go h.readPump()

	h.connected = true
	h.mu.Unlock()
	return nil
}

func (h *WebSocketHandler) connectWebSocket() error {
	dialer := websocket.Dialer{}
	var err error

	h.conn, _, err = dialer.Dial(h.WS_URL, h.headers)
	if err != nil {
		log.Info("Failed to connect", "ws_url", h.WS_URL)
		return fmt.Errorf("failed to connect to websocket: %v", err)
	}
	return nil
}

func (h *WebSocketHandler) readPump() {
	defer func() {
		h.mu.Lock()
		if h.conn != nil {
			h.conn.Close()
		}
		h.connected = false
		h.mu.Unlock()

		// Try to reconnect after a delay
		// go func() {
		// 	select {
		// 	case <-h.done:
		// 		return
		// 	case <-time.After(5 * time.Second):
		// 		err := h.Connect()
		// 		if err != nil {
		// 			log.Printf("Reconnection failed: %v", err)
		// 		}
		// 	}
		// }()
	}()

	for {
		select {
		case <-h.done:
			return
		default:
			if !h.connected {
				return
			}

			_, p, err := h.conn.ReadMessage()
			log.Info("Got message", "msg", string(p))
			if err != nil {
        log.Error("Error in readPump()", "err", err)
				if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					log.Printf("Error reading message: %v", err)
				}
				return
			}
			h.msgChan <- p
      log.Info("Passed message to upper handler")
		}
	}
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
	h.mu.Lock()
	defer h.mu.Unlock()

	if !h.connected {
		return
	}

	close(h.done)
	if h.conn != nil {
		h.conn.Close()
	}
	h.connected = false
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
	log.Info("FocusNow() returns")

	return nil
}
