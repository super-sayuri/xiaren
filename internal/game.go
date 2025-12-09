package internal

import (
	"context"
	"xiaren/internal/layout"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	W  int
	H  int
	Ui *ebitenui.UI
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
	defaultMainMenu := &layout.DefaultMainMenu{
		W: w,
		H: h,
	}
	return &Game{
		W: w,
		H: h,
		Ui: &ebitenui.UI{
			Container: defaultMainMenu.GetContainer(context.Background()),
		},
	}
}
