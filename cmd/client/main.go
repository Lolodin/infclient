package main

import (
	"github.com/Lolodin/infclient/internal/interactive"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/text/language"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 700
	screenHeight = 590
)
const MainMenu = "main"

func main() {
	game := kernel.NewState(kernel.Options{
		Title:        "test",
		RenderHeight: screenHeight,
		RenderWidth:  screenWidth,
	}, language.Russian)

	game.LoadWorld(interactive.NewTitleWorld(game))
	game.ActivateWorlds(MainMenu)
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatalln(err)
	}
}
