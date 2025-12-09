package render

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderLayer interface {
	Draw(*ebiten.Image)
	Update(context.Context, any) error
	ZIndex() int
}

type LayerManager struct {
	layers []RenderLayer
}

func NewLayerManager() *LayerManager {
	return &LayerManager{
		layers: make([]RenderLayer, 0),
	}
}

func (lm *LayerManager) AddLayer(layer RenderLayer) {
	lm.layers = append(lm.layers, layer)
	lm.sort()
}

func (lm *LayerManager) RemoveLayer(layer RenderLayer) {
	for i, l := range lm.layers {
		if l == layer {
			lm.layers = append(lm.layers[:i], lm.layers[i+1:]...)
			break
		}
	}
}

func (lm *LayerManager) sort() {
	// 按 ZIndex 排序
	for i := 0; i < len(lm.layers)-1; i++ {
		for j := i + 1; j < len(lm.layers); j++ {
			if lm.layers[i].ZIndex() > lm.layers[j].ZIndex() {
				lm.layers[i], lm.layers[j] = lm.layers[j], lm.layers[i]
			}
		}
	}
}

func (lm *LayerManager) Draw(screen *ebiten.Image) {
	for _, layer := range lm.layers {
		layer.Draw(screen)
	}
}
