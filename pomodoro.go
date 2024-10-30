package main

import (
	"fmt"
	"log"
	"time"
)

type PomodoroState string

const (
	StateIdle     PomodoroState = "idle"
	StateRunning  PomodoroState = "running"
	StatePaused   PomodoroState = "paused"
	StateFinished PomodoroState = "finished"
)

type Pomodoro struct {
	Duration     time.Duration
	TimeLeft     time.Duration
	State        PomodoroState
	StartTime    time.Time
	timer        *time.Timer
	ticker       *time.Ticker
	tickCallback func(float64)
}

func NewPomodoro(duration time.Duration) *Pomodoro {
	return &Pomodoro{
		Duration: duration,
		TimeLeft: duration,
		State:    StateIdle,
	}
}

func (p *Pomodoro) Start() {
	if p.State == StateRunning {
		return
	}

	p.State = StateRunning
	p.StartTime = time.Now()
	p.timer = time.NewTimer(p.TimeLeft)
	p.ticker = time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-p.timer.C:
				p.State = StateFinished
				p.TimeLeft = 0
				if p.ticker != nil {
					p.ticker.Stop()
					p.ticker = nil
				}
				if p.tickCallback != nil {
					p.tickCallback(0)
				}
				return
			case <-p.ticker.C:
				p.TimeLeft = p.TimeLeft - time.Second
        log.Println("tick")
				if p.tickCallback != nil {
					p.tickCallback(p.TimeLeft.Seconds())
				}
			}
		}
	}()
}

func (p *Pomodoro) Stop() {
	if p.State != StateRunning {
		return
	}

	if !p.timer.Stop() {
		select {
		case <-p.timer.C:
		default:
		}
	}

	if p.ticker != nil {
		p.ticker.Stop()
		p.ticker = nil
	}

	p.State = StateIdle
	p.TimeLeft = p.Duration
	if p.tickCallback != nil {
		p.tickCallback(p.Duration.Seconds())
	}
}

func (p *Pomodoro) Pause() {
	if p.State != StateRunning {
		return
	}

	if !p.timer.Stop() {
		select {
		case <-p.timer.C:
		default:
		}
	}

	p.TimeLeft = p.TimeLeft - time.Since(p.StartTime)
	p.State = StatePaused
}

func (p *Pomodoro) Resume() {
	if p.State != StatePaused {
		return
	}

	p.Start()
}

func (p *Pomodoro) SetTickCallback(callback func(float64)) {
	p.tickCallback = callback
}
