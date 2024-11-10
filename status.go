package main

import "time"

type Status struct {
	Focusing  bool
	TimeSince time.Duration
	done      chan struct{}
}

func NewStatus() *Status {
	s := &Status{
		Focusing:  false,
		TimeSince: 0,
		done:      make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.TimeSince += time.Minute
			case <-s.done:
				return
			}
		}
	}()

	return s
}

func (s *Status) Close() {
	close(s.done)
}

func (s *Status) SetFocusing(focusing bool) {
	s.Focusing = focusing
	s.TimeSince = 0
}
