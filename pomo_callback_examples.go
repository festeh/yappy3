package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"yappy3/astal"
	"yappy3/pomodoro"
)

func FormatTime(seconds float64) string {
	minutes := int(seconds) / 60
	secs := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d", minutes, secs)
}

func TickTimeLeftWrapper(ctx context.Context, p *pomodoro.Pomo) func(p *pomodoro.Pomo) {
	return func(p *pomodoro.Pomo) {
		timeLeft := FormatTime(p.TimeLeft.Seconds())
		runtime.EventsEmit(ctx, "tick", timeLeft)
	}
}

func TickTimeLeftAstal(p *pomodoro.Pomo) {
	astal := astal.Astal{}
	// If more than one minute, show rounded to nearest minutes
	// show seconds otherwise
	timeLeft := fmt.Sprintf("%d", int(p.TimeLeft.Seconds()/60))
	if p.TimeLeft.Seconds() < 60 {
		timeLeft = fmt.Sprintf("%d", int(p.TimeLeft.Seconds()))
	}
	astal.SendMessage(fmt.Sprintf("{\"pomodoro\": \"%s\"}", timeLeft))
}

func StopResetTimeWrapper(ctx context.Context, p *pomodoro.Pomo) func(p *pomodoro.Pomo) {
	return func(p *pomodoro.Pomo) {
		duration := FormatTime(p.Duration.Seconds())
		runtime.EventsEmit(ctx, "tick", duration)
	}
}
