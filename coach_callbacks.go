package main

import (
	"context"
	"fmt"
	"yappy3/astal"
	"yappy3/coach"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EmitOnFocusSetWrapper(ctx context.Context, c *coach.Coach) func(*coach.Coach) {
	return func(c *coach.Coach) {
		runtime.EventsEmit(ctx, "focusing", c.Focusing)
	}
}

func OnFocusSetAstal(c *coach.Coach) {
	astal := astal.Astal{}
  msg := "Not focusing :("
  if c.Focusing {
    msg = "Focusing!"
  }
	astal.SendMessage(fmt.Sprintf("{\"focusing\":\"%s\"}", msg))
}
