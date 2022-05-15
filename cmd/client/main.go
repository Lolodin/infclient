package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Endor")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
