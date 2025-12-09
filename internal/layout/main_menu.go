package layout

import (
	"context"
	"xiaren/internal/resource"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/inconsolata"
)

type DefaultMainMenu struct {
	W int
	H int
}

func (m *DefaultMainMenu) GetContainer(ctx context.Context) *widget.Container {
	paddingUp := 0.2
	paddingDown := 0.2
	paddingLeft := 0.15
	bgImage, err := resource.LoadImage(ctx, "backgrounds", "bg1.png")
	if err != nil {
		panic(err)
	}
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			//image.NewNineSliceColor(colornames.Mistyrose),
			image.NewFixedNineSlice(bgImage),
		),
		widget.ContainerOpts.Layout(widget.NewStackedLayout(
			widget.StackedLayoutOpts.Padding(
				&widget.Insets{
					Top:    int(paddingUp * float64(m.H)),
					Bottom: int(paddingDown * float64(m.H)),
				},
			),
		)),
	)

	menu := widget.NewContainer(
		// widget.ContainerOpts.BackgroundImage(
		// 	image.NewNineSliceColor(colornames.Orange),
		// ),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(
				&widget.Insets{
					Top:  1,
					Left: int(paddingLeft * float64(m.W)),
				},
			),
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(int(0.0625*float64(m.H))),
		)),
	)
	menu.AddChild(m.mainMenuButton("New Game"))
	menu.AddChild(m.mainMenuButton("读取游戏"))
	menu.AddChild(m.mainMenuButton("设置"))
	menu.AddChild(m.mainMenuButton("画廊"))
	menu.AddChild(m.mainMenuButton("退出至桌面"))
	root.AddChild(menu)
	return root
}

func (m *DefaultMainMenu) mainMenuButton(label string) *widget.Button {
	var defaultFont text.Face = text.NewGoXFace(inconsolata.Regular8x16)
	return widget.NewButton(
		widget.ButtonOpts.TextLabel(label),
		widget.ButtonOpts.TextFace(&defaultFont),
		widget.ButtonOpts.TextColor(
			&widget.ButtonTextColor{
				Idle:    colornames.White,
				Hover:   colornames.Black,
				Pressed: colornames.Gray,
			},
		),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    image.NewNineSliceColor(colornames.Darkslategray),
			Hover:   image.NewNineSliceColor(colornames.Mediumseagreen),
			Pressed: image.NewNineSliceColor(colornames.Black),
		}),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(int(0.141*float64(m.W)), int(0.07*float64(m.H))),
		),
	)
}
