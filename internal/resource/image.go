package resource

import (
	"context"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func LoadImage(ctx context.Context, entry ...string) (*ebiten.Image, error) {
	fullEntry := filepath.Join(append([]string{"images"}, entry...)...)
	resource := GetResource()
	reader, err := resource.GetAsset(ctx, fullEntry)
	if err != nil {
		return nil, err
	}
	ei, _, err := ebitenutil.NewImageFromReader(reader)
	return ei, err
}
