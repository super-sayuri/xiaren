package layout

import (
	"context"
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"xiaren/internal/render"
)

type DefaultDialogue struct {
	BaseLayout
}

func (d *DefaultDialogue) Draw(screen *ebiten.Image) {
	ui := &ebitenui.UI{}
	ui.Draw(screen)
}

func (d *DefaultDialogue) Update(_ context.Context, _ any) error {
	return nil
}

func (d *DefaultDialogue) ZIndex() int {
	return render.ZIndexDialogue
}

var _ render.RenderLayer = (*DefaultDialogue)(nil)
