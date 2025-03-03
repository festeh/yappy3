package coach

import (
	"context"
	"fmt"
	"yappy3/astal"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func EmitOnFocusSetWrapper(ctx context.Context, c *Coach) func(*Coach) {
	return func(c *Coach) {
		runtime.EventsEmit(ctx, "focusing", c.Focusing)
	}
}

func OnFocusSetAstal(c *Coach) {
	astal := astal.Astal{}
	msg := "Not focusing :("
	if c.Focusing {
		msg = "Focusing!"
	}
	astal.SendMessage(fmt.Sprintf("{\"focusing\":\"%s\"}", msg))
}
