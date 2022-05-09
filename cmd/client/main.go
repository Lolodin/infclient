package main

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/object"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	Login object.Objecter
}

func NewGame() *Game {
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
	obj := object.NewObject(text.BoundString(mplusNormalFont, "Login"), screenWidth/2.5, screenHeight/3)
	obj.Font = mplusNormalFont
	obj.Clr = color.White
	obj.Label = "Login"
	return &Game{Login: obj}
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

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Endor")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
