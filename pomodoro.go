package main

import (
	"fmt"
	"log"
	"time"
)

func formatTime(seconds float64) string {
	minutes := int(seconds) / 60
	secs := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

type ButtonInfo struct {
	Text   string `json:"text"`
	Method string `json:"method"`
}

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
	tickCallback func(string)
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
					p.tickCallback(formatTime(0))
				}
				return
			case <-p.ticker.C:
				p.TimeLeft = p.TimeLeft - time.Second
				log.Println("tick")
				if p.tickCallback != nil {
					p.tickCallback(formatTime(p.TimeLeft.Seconds()))
				}
			}
		}
	}()
}

func (p *Pomodoro) Stop() {
	log.Println("Stopping")
	if p.State == StateIdle || p.State == StateFinished {
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
		log.Println("Stopped ticker")
		p.ticker = nil
	}

	p.State = StateIdle
	p.TimeLeft = p.Duration
	if p.tickCallback != nil {
		p.tickCallback(formatTime(p.Duration.Seconds()))
	}
}

func (p *Pomodoro) Pause() {
	log.Println("Paused")
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

	// p.TimeLeft = p.TimeLeft - time.Since(p.StartTime)
	p.State = StatePaused
}

func (p *Pomodoro) Resume() {
	if p.State != StatePaused {
		return
	}

	p.Start()
}

func (p *Pomodoro) SetTickCallback(callback func(string)) {
	p.tickCallback = callback
}

func (p *Pomodoro) GetButtons() []ButtonInfo {
	switch p.State {
	case StateIdle:
		return []ButtonInfo{{Text: "Start", Method: "StartPomodoro"}}
	case StateRunning:
		return []ButtonInfo{
			{Text: "Pause", Method: "PausePomodoro"},
			{Text: "Stop", Method: "StopPomodoro"},
		}
	case StatePaused:
		return []ButtonInfo{
			{Text: "Resume", Method: "ResumePomodoro"},
			{Text: "Stop", Method: "StopPomodoro"},
		}
	case StateFinished:
		return []ButtonInfo{{Text: "Start New", Method: "StartPomodoro"}}
	default:
		return []ButtonInfo{}
	}
}
