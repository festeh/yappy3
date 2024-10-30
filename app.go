package main

import (
	"context"
	"time"

	"github.com/wailsjs/runtime"
)

// App struct
type App struct {
	ctx      context.Context
	pomodoro *Pomodoro
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		pomodoro: NewPomodoro(25 * time.Minute),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.pomodoro.SetTickCallback(func(timeLeft float64) {
		runtime.EventsEmit(ctx, "pomodoroTick", timeLeft)
	})
}

// StartPomodoro starts the pomodoro timer
func (a *App) StartPomodoro() {
	a.pomodoro.Start()
}

// StopPomodoro stops the pomodoro timer
func (a *App) StopPomodoro() {
	a.pomodoro.Stop()
}

// PausePomodoro pauses the pomodoro timer
func (a *App) PausePomodoro() {
	a.pomodoro.Pause()
}

// ResumePomodoro resumes the paused pomodoro timer
func (a *App) ResumePomodoro() {
	a.pomodoro.Resume()
}

// GetPomodoroState returns the current state of the pomodoro timer
func (a *App) GetPomodoroState() PomodoroState {
	return a.pomodoro.State
}

// GetTimeLeft returns the remaining time in seconds
func (a *App) GetTimeLeft() float64 {
	return a.pomodoro.TimeLeft.Seconds()
}

// SetPomodoroTime sets a new duration for the pomodoro timer (in minutes)
func (a *App) SetPomodoroTime(minutes float64) {
	newPomodoro := NewPomodoro(time.Duration(minutes * float64(time.Minute)))
	newPomodoro.SetTickCallback(func(timeLeft float64) {
		runtime.EventsEmit(a.ctx, "pomodoroTick", timeLeft)
	})
	a.pomodoro = newPomodoro
}
