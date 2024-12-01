package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
  "yappy3/pomodoro"
)

func FormatTime(seconds float64) string {
	minutes := int(seconds) / 60
	secs := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

func TimeLeftOnTickWrapper(ctx context.Context, p *pomodoro.Pomo) func(p *pomodoro.Pomo) {
	return func(p *pomodoro.Pomo) {
		timeLeft := FormatTime(p.TimeLeft.Seconds())
		runtime.EventsEmit(ctx, "tick", timeLeft)
	}
}

func ResetTimeOnStopWrapper(ctx context.Context, p *pomodoro.Pomo) func(p *pomodoro.Pomo) {
  return func(p *pomodoro.Pomo) {
    duration := FormatTime(p.Duration.Seconds())
    runtime.EventsEmit(ctx, "tick", duration)
  }
}
