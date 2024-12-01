package main

import (
	"time"
	"yappy3/pomodoro"
)

// StartPomodoro starts the pomodoro timer
func (a *App) StartPomodoro() {
	a.pomo.Start()
}

// StopPomodoro stops the pomodoro timer
func (a *App) StopPomodoro() {
	a.pomo.Stop()
}

// PausePomodoro pauses the pomodoro timer
func (a *App) PausePomodoro() {
	a.pomo.Pause()
}

// ResumePomodoro resumes the paused pomodoro timer
func (a *App) ResumePomodoro() {
	a.pomo.Resume()
}

// GetPomodoroState returns the current state of the pomodoro timer
func (a *App) GetPomodoroState() pomodoro.PomodoroState {
	return a.pomo.State
}

// GetTimeLeft returns the remaining time in MM:SS format
func (a *App) GetTimeLeft() string {
	return FormatTime(a.pomo.TimeLeft.Seconds())
}

// SetPomodoroTime sets a new duration for the pomodoro timer (in minutes)
func (a *App) SetPomodoroTime(minutes float64) {
	newPomodoro := pomodoro.NewPomodoro(time.Duration(minutes * float64(time.Minute)))
	a.pomo = newPomodoro
}

// GetPomodoroButtons returns the available buttons based on current state
func (a *App) GetPomodoroButtons() []pomodoro.ButtonInfo {
	return a.pomo.GetButtons()
}
