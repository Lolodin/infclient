package kernel

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
}

func NewGame() *Game {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	text.Draw(screen, g.Login.GetLabel(), g.Login.GetFont(), g.Login.GetX(), g.Login.GetY(), g.Login.GetColor())
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.Login.In(ebiten.CursorPosition()) {
			fmt.Print("click")
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
