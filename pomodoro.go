package main

import (
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

	go func() {
		<-p.timer.C
		p.State = StateFinished
		p.TimeLeft = 0
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

	p.State = StateIdle
	p.TimeLeft = p.Duration
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
