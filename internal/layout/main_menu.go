package layout

import (
	"context"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/inconsolata"
	"os"
	"xiaren/internal/eventbus"
	"xiaren/internal/log"
	"xiaren/internal/render"
	"xiaren/internal/resource"
)

type DefaultMainMenu struct {
	BaseLayout
	Background *image.NineSlice
	TextFace   text.Face
}

type DefaultMainMenuOpt func(*DefaultMainMenu)

var DefaultMainMenuBaseOpt = func(base BaseLayout) DefaultMainMenuOpt {
	return func(m *DefaultMainMenu) {
		m.BaseLayout = base
	}
}

var DefaultMainMenuBackgroundOpt = func(bg *image.NineSlice) DefaultMainMenuOpt {
	return func(m *DefaultMainMenu) {
		m.Background = bg
	}
}

var DefaultMainMenuTextFaceOpt = func(face text.Face) DefaultMainMenuOpt {
	return func(m *DefaultMainMenu) {
		m.TextFace = face
	}
}

func NewDefaultMainMenu(ctx context.Context, opts ...DefaultMainMenuOpt) *DefaultMainMenu {
	logger := log.GetLog(ctx)
	logger.Debug("Loading main menu resources...")
	defer logger.Debug("Main menu resources loaded.")

	backgroundImage := image.NewNineSliceColor(colornames.Magenta)
	img, err := resource.LoadImage(ctx, "backgrounds", "bg1.png")
	if err != nil {
		logger.Info("Failed to load main menu background image, using solid color background instead:", err)
	} else {
		backgroundImage = image.NewFixedNineSlice(img)
	}

	textFace := text.NewGoXFace(inconsolata.Regular8x16)

	menu := &DefaultMainMenu{
		BaseLayout: BaseLayout{},
		Background: backgroundImage,
		TextFace:   textFace,
	}
	for _, opt := range opts {
		opt(menu)
	}
	return menu
}

func (m *DefaultMainMenu) GetContainer(ctx context.Context) *widget.Container {
	paddingUp := 0.2
	paddingDown := 0.2
	paddingLeft := 0.15
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			m.Background,
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
	menu.AddChild(m.mainMenuButton("New Game",
		func(_ *widget.ButtonClickedEventArgs) {
			m.EventBus.Publish(eventbus.CHANGE_LAYOUT_TOPIC, "dialogue")
		}))
	menu.AddChild(m.mainMenuButton("读取游戏",
		func(_ *widget.ButtonClickedEventArgs) {
			m.EventBus.Publish(eventbus.CHANGE_LAYOUT_TOPIC, "story")
		}))
	menu.AddChild(m.mainMenuButton("设置",
		func(_ *widget.ButtonClickedEventArgs) {
			m.EventBus.Publish(eventbus.CHANGE_LAYOUT_TOPIC, "story")
		}))
	menu.AddChild(m.mainMenuButton("画廊",
		func(_ *widget.ButtonClickedEventArgs) {
			m.EventBus.Publish(eventbus.CHANGE_LAYOUT_TOPIC, "story")
		}))
	menu.AddChild(m.mainMenuButton("退出至桌面",
		func(_ *widget.ButtonClickedEventArgs) {
			os.Exit(0)
		}))
	root.AddChild(menu)
	return root
}

func (m *DefaultMainMenu) mainMenuButton(label string, f widget.ButtonClickedHandlerFunc) *widget.Button {
	return widget.NewButton(
		widget.ButtonOpts.TextLabel(label),
		widget.ButtonOpts.TextFace(&m.TextFace),
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
		widget.ButtonOpts.ClickedHandler(f),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(int(0.141*float64(m.W)), int(0.07*float64(m.H))),
		),
	)
}

func (m *DefaultMainMenu) Update(_ context.Context, _ any) error {
	return nil
}

func (m *DefaultMainMenu) Draw(screen *ebiten.Image) {
	ui := &ebitenui.UI{
		Container: m.GetContainer(context.Background()),
	}
	ui.Draw(screen)
}

func (m *DefaultMainMenu) ZIndex() int {
	return render.ZIndexBackground
}

var _ render.RenderLayer = (*DefaultMainMenu)(nil)
