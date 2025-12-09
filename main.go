package main

import (
	"context"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	xia "xiaren/internal"
)

const DEFAULT_RESOURCE_PATH = "assets/default"

func main() {
	var err error

	err = xia.Setup(context.Background(), DEFAULT_RESOURCE_PATH)
	if err != nil {
		panic(err)
	}

	resolutionX := xia.GetSetting().WindowSetting.ResolutionX
	resolutionY := xia.GetSetting().WindowSetting.ResolutionY
	ebiten.SetWindowSize(resolutionX, resolutionY)
	ebiten.SetWindowTitle(xia.GetSetting().GameSetting.Title)
	if err := ebiten.RunGame(xia.NewGame(resolutionX, resolutionY)); err != nil {
		log.Fatal(err)
	}
}
