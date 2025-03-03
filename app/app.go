package app

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
	"yappy3/coach"
	"yappy3/pomodoro"
)

func onReady() {
	systray.SetTitle("yappy3")
	systray.SetTooltip("yappy3")
	systray.AddMenuItem("Quit", "Quit the whole app")
	systray.SetIcon(icon)
}

func onExit() {
	systray.Quit()
}

// App struct
type App struct {
	ctx   context.Context
	Pomo  *pomodoro.Pomo
	Coach *coach.Coach
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
		Pomo:  pomodoro.NewPomodoro(25 * time.Minute),
		Coach: coach.NewCoach(wsurl, url),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	pCbs := a.Pomo.Callbacks
	pCbs.AddTick(pomodoro.TickTimeLeftWrapper(ctx, a.Pomo))
	pCbs.AddTick(pomodoro.TickTimeLeftAstal)

	pCbs.AddStart(pomodoro.NotifyPomodoroStart)
	pCbs.AddStop(pomodoro.NotifyPomodoroStop)
	pCbs.AddStop(pomodoro.StopResetTimeWrapper(ctx, a.Pomo))
	pCbs.AddFinish(pomodoro.NotifyPomodoroFinish)

	cCbs := a.Coach.Callbacks
	cCbs.OnFocusReceived = append(cCbs.OnFocusReceived, coach.EmitOnFocusSetWrapper(ctx, a.Coach))
	cCbs.OnFocusReceived = append(cCbs.OnFocusReceived, coach.OnFocusSetAstal)

	go func() {
		// Wait for Wails to fully initialize
		time.Sleep(500 * time.Millisecond)
		systray.Run(onReady, onExit)
	}()
}

// SetPomodoroTime sets a new duration for the pomodoro timer (in minutes)
func (a *App) SetPomodoroTime(minutes float64) {
	newPomodoro := pomodoro.NewPomodoro(time.Duration(minutes * float64(time.Minute)))
	a.Pomo = newPomodoro
}
