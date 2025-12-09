package internal

import (
	"context"
	"xiaren/internal/eventbus"
	"xiaren/internal/layout"
	"xiaren/internal/render"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	W             int
	H             int
	eventBus      *eventbus.EventBus
	renderManager render.LayerManager
	Ui            *ebitenui.UI
}

func (g *Game) Update() error {
	g.Ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

var _ ebiten.Game = (*Game)(nil)

func NewGame(w, h int) *Game {
	eventBus := eventbus.NewEventBus()
	defaultMainMenu := &layout.DefaultMainMenu{
		BaseLayout: layout.BaseLayout{
			W:        w,
			H:        h,
			EventBus: eventBus,
		},
	}
	layoutCh := make(chan any)
	eventBus.Subscribe(eventbus.CHANGE_LAYOUT_TOPIC, layoutCh)
	game := &Game{
		W: w,
		H: h,
		Ui: &ebitenui.UI{
			Container: defaultMainMenu.GetContainer(context.Background()),
		},
		eventBus: eventBus,
	}
	go func(g *Game) {
		for {
			select {
			case ev := <-layoutCh:
				panic("todo: layout change not implemented: " + ev.(string))
			}
		}
	}(game)
	return game
}
