package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
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
	pCbs := a.pomo.Callbacks
	pCbs.AddTick(pomodoro.TickTimeLeftWrapper(ctx, a.pomo))
	pCbs.AddTick(pomodoro.TickTimeLeftAstal)

	pCbs.AddStart(pomodoro.NotifyPomodoroStart)
	pCbs.AddStop(pomodoro.NotifyPomodoroStop)
	pCbs.AddStop(pomodoro.StopResetTimeWrapper(ctx, a.pomo))
	pCbs.AddFinish(pomodoro.NotifyPomodoroFinish)

	cCbs := a.coach.Callbacks
	cCbs.OnFocusReceived = append(cCbs.OnFocusReceived, coach.EmitOnFocusSetWrapper(ctx, a.coach))
	cCbs.OnFocusReceived = append(cCbs.OnFocusReceived, coach.OnFocusSetAstal)
}
