package coach

import "time"

type Coach struct {
	Focusing   bool
	TimeSince  time.Duration
	handler    *WebSocketHandler
	onFocusSet func(focusing bool)
}

func NewCoach(ws_url string, url string) *Coach {
	s := &Coach{
		Focusing:   false,
		TimeSince:  0,
		handler:    NewWebSocketHandler(ws_url, url),
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.TimeSince += time.Minute
			case <-s.handler.done:
				return
			}
		}
	}()

	return s
}

func (s *Coach) Close() {
	close(s.handler.done)
}

func (s *Coach) SetFocusing(focusing bool) {
	s.Focusing = focusing
	s.TimeSince = 0
	s.onFocusSet(focusing)
}

func (s *Coach) SetOnFocusSet(f func(focusing bool)) {
  s.onFocusSet = f
}
