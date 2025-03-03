package pomodoro

import (
	"log"
	"time"
)

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

type Pomo struct {
	Duration  time.Duration
	TimeLeft  time.Duration
	State     PomodoroState
	StartTime time.Time
	timer     *time.Timer
	ticker    *time.Ticker
	Callbacks *Callbacks
}

func NewPomodoro(duration time.Duration) *Pomo {
	return &Pomo{
		Duration:  duration,
		TimeLeft:  duration,
		State:     StateIdle,
		Callbacks: NewCallbacks(),
	}
}

func (p *Pomo) Start() {
	if p.State == StateRunning {
		return
	}

	p.State = StateRunning
	p.StartTime = time.Now()
	p.timer = time.NewTimer(p.TimeLeft)
	p.ticker = time.NewTicker(time.Second)

	p.Callbacks.RunOnStart(p)

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
				p.Callbacks.RunOnFinish(p)
				return
			case <-p.ticker.C:
				p.TimeLeft = p.TimeLeft - time.Second
				p.Callbacks.RunOnTick(p)
			}
		}
	}()
}

func (p *Pomo) Stop() {
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

	p.Callbacks.RunOnStop(p)
}

func (p *Pomo) Pause() {
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

	p.State = StatePaused
}

func (p *Pomo) Resume() {
	if p.State != StatePaused {
		return
	}

	p.Start()
}

// GetTimeLeft returns the remaining time in MM:SS format
func (p *Pomo) GetTimeLeft() string {
  return FormatTime(p.TimeLeft.Seconds())
}



func (p *Pomo) GetButtons() []ButtonInfo {
	switch p.State {
	case StateIdle:
		return []ButtonInfo{{Text: "Start", Method: "start"}}
	case StateRunning:
		return []ButtonInfo{
			{Text: "Pause", Method: "pause"},
			{Text: "Stop", Method: "stop"},
		}
	case StatePaused:
		return []ButtonInfo{
			{Text: "Resume", Method: "resume"},
			{Text: "Stop", Method: "stop"},
		}
	case StateFinished:
		return []ButtonInfo{{Text: "Start New", Method: "start"}}
	default:
		return []ButtonInfo{}
	}
}
