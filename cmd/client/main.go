package main

import (
	"fmt"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	runnerImage *ebiten.Image
)

type Game struct {
	Login image.Rectangle
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	var mplusNormalFont font.Face
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     78,
		Hinting: font.HintingFull,
	})

	g.Login = text.BoundString(mplusNormalFont, "Login")
	text.Draw(screen, "Login", mplusNormalFont, screenWidth/2, screenHeight/3, color.White)
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		if g.Login.RGBA64At(x-screenWidth/2, y-screenHeight/3).A > 0 {
			fmt.Print("click")
		}

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Endor")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
