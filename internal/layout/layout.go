package layout

import (
	"xiaren/internal/eventbus"
	"xiaren/internal/render"
)

type BaseLayout struct {
	W        int
	H        int
	EventBus *eventbus.EventBus
}

func LayoutFactory(name string, args BaseLayout) render.RenderLayer {
	switch name {
	case "default_main_menu":
		return &DefaultMainMenu{
			BaseLayout: args,
		}
	default:
		return &DefaultMainMenu{
			BaseLayout: args,
		}
	}
}
