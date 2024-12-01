package coach

import (
	"encoding/json"
	"log"
	"time"
)

type Coach struct {
	Focusing  bool
	TimeSince time.Duration
	handler   *WebSocketHandler
	callbacks *Callbacks
}

func NewCoach(ws_url string, url string) *Coach {
	s := &Coach{
		Focusing:  false,
		TimeSince: 0,
		handler:   NewWebSocketHandler(ws_url, url),
		callbacks: NewCallbacks(),
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.TimeSince += time.Minute
				s.callbacks.RunOnTick(s)
			case <-s.handler.done:
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
	return s.Focusing
}

func (s *Coach) Close() {
	close(s.handler.done)
}

func (s *Coach) SetFocusing(focusing bool) {
	log.Printf("Focus state updated: %v", s.Focusing)
	s.Focusing = focusing
	s.TimeSince = 0
	s.onFocusSet(focusing)
}

func (s *Coach) SetOnFocusSet(f func(focusing bool)) {
	s.onFocusSet = f
}

func (s *Coach) handleFocus(message []byte) {
	// Parse the message to get focus state
	var msg struct {
		Event    string `json:"event"`
		Focusing bool   `json:"focusing"`
	}
	if err := json.Unmarshal(message, &msg); err != nil {
		log.Printf("Error parsing focus message: %v", err)
		return
	}

	s.SetFocusing(msg.Focusing)
}
