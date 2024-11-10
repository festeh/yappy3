package main

import "time"

type Status struct {
	Focusing  bool
	TimeSince time.Duration
}

func NewStatus() *Status {
	return &Status{
		Focusing:  false,
		TimeSince: 0,
	}
}

func (s *Status) SetFocusing(focusing bool) {
	s.Focusing = focusing
	s.TimeSince = 0
}
