package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"yappy3/coach"
	"yappy3/pomodoro"
)

// App struct
type App struct {
	ctx   context.Context
	pomo  *pomodoro.Pomo
	coach *coach.Coach
}

// NewApp creates a new App application struct
func NewApp() *App {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using defaults")
	}

	wsurl := os.Getenv("COACH_WSURL")
	if wsurl == "" {
		wsurl = "ws://localhost:8080" // fallback default
		log.Printf("WSURL environment variable not set, using default: %s", wsurl)
	}

	url := os.Getenv("COACH_URL")
	if url == "" {
		url = "http://localhost:8080" // fallback default
		log.Printf("URL environment variable not set, using default: %s", url)
	}

	return &App{
		pomo:  pomodoro.NewPomodoro(25 * time.Minute),
		coach: coach.NewCoach(wsurl, url),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.pomo.SetTickCallback(func(timeLeft string) {
		runtime.EventsEmit(ctx, "tick", timeLeft)
	})
	a.coach.SetOnFocusSet(func(focusing bool) {
		runtime.EventsEmit(ctx, "focusing", focusing)
	})
}

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
	return pomodoro.FormatTime(a.pomo.TimeLeft.Seconds())
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
