package pomodoro

import (
	"context"
	"fmt"
	"math"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"yappy3/astal"
)

func FormatTime(seconds float64) string {
	minutes := int(seconds) / 60
	secs := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

func TickTimeLeftWrapper(ctx context.Context, p *Pomo) func(p *Pomo) {
	return func(p *Pomo) {
		timeLeft := FormatTime(p.TimeLeft.Seconds())
		runtime.EventsEmit(ctx, "tick", timeLeft)
	}
}

func TickTimeLeftAstal(p *Pomo) {
	astal := astal.Astal{}
	seconds := p.TimeLeft.Seconds()
	var timeLeft string

	if seconds >= 60 {
		minutes := math.Round(seconds / 60)
		suffix := "mins"
		if minutes == 1 {
			suffix = "min"
		}
		timeLeft = fmt.Sprintf("%d %s", int(minutes), suffix)
	} else {
		suffix := "secs"
		if seconds == 1 {
			suffix = "sec"
		}
		timeLeft = fmt.Sprintf("%d %s", int(seconds), suffix)
	}
	astal.SendMessage(fmt.Sprintf("{\"pomodoro\": \"%s\"}", timeLeft))
}

func NotifyPomodoroStart(p *Pomo) {
	cmd := exec.Command("notify-send", "Pomodoro", "Pomodoro has started")
	cmd.Run()
}

func NotifyPomodoroStop(p *Pomo) {
	cmd := exec.Command("notify-send", "Pomodoro", "Pomodoro has been stopped")
	cmd.Run()
}

func NotifyPomodoroFinish(p *Pomo) {
	cmd := exec.Command("notify-send", "Pomodoro", "Pomodoro finished! Hooray!")
	cmd.Run()
}

func StopResetTimeWrapper(ctx context.Context, p *Pomo) func(p *Pomo) {
	return func(p *Pomo) {
		duration := FormatTime(p.Duration.Seconds())
		runtime.EventsEmit(ctx, "tick", duration)
	}
}
