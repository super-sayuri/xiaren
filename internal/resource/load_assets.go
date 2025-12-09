package resource

import (
	"context"
	"io"
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

func LoadFont(ctx context.Context, entry ...string) ([]byte, error) {
	fullEntry := filepath.Join(append([]string{"fonts"}, entry...)...)
	resource := GetResource()
	reader, err := resource.GetAsset(ctx, fullEntry)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(reader)
	return data, err
}
