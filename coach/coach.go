package coach

import (
	"encoding/json"
  "github.com/charmbracelet/log"
	"time"
)

type Coach struct {
	Focusing  bool
	TimeSince time.Duration
	handler   *WebSocketHandler
	Callbacks *Callbacks
}

func NewCoach(ws_url string, url string) *Coach {
	s := &Coach{
		Focusing:  false,
		TimeSince: 0,
		handler:   NewWebSocketHandler(ws_url, url),
		Callbacks: NewCallbacks(),
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
        log.Info("Tick")
				s.TimeSince += time.Minute
				s.Callbacks.RunOnTick(s)
			case msg := <-s.handler.msgChan:
        log.Info("Got message")
				s.handleMessage(msg)
			case <-s.handler.done:
        log.Info("We're done")
				return
			}
		}
	}()

	return s
}

func (s *Coach) Connect() {
	s.handler.Connect()
}

func (s *Coach) Disconnect() {
	s.handler.Disconnect()
}

func (s *Coach) GetFocusing() bool {
	res, err := s.handler.GetFocus()
	if err != nil {
		log.Error("Error getting focus", "err", err)
		return false
	}
	return res
}

func (s *Coach) Close() {
	close(s.handler.done)
}

func (s *Coach) SetFocusing(focusing bool) {
	log.Info("SetFocusing()", "focusing", focusing)
	s.Focusing = focusing
	s.TimeSince = 0
	s.Callbacks.RunOnFocusReceived(s)
}

func (s *Coach) handleMessage(message []byte) {
  log.Info("Handling message", "message", string(message))
	// Parse JSON message
	var msg struct {
		Event    string `json:"event"`
		Focusing bool   `json:"focusing"`
	}
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Printf("Error parsing message: %v", err)
		return
	}

	// Handle different event types
	switch msg.Event {
	case "focusing":
		s.SetFocusing(msg.Focusing)
	default:
		log.Printf("Unknown event: %s", msg.Event)
	}
}

func (s *Coach) FocusNow() {
  s.handler.FocusNow()
}
