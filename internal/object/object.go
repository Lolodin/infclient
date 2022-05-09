package object

import (
	"golang.org/x/image/font"
	"image"
	"image/color"
)

type Objecter interface {
	In(x, y int) bool
	GetLabel() string
	GetFont() font.Face
	GetColor() color.Color
	GetX() int
	GetY() int
}

type Object struct {
	Image image.Rectangle
	Clr   color.Color
	Font  font.Face
	Label string
	X     int
	Y     int
}

func (o Object) GetColor() color.Color {
	return o.Clr
}

func (o Object) GetX() int {
	return o.X
}

func (o Object) GetY() int {
	return o.Y
}

func NewObject(image image.Rectangle, x int, y int) *Object {
	return &Object{Image: image, X: x, Y: y}
}

func (o Object) GetLabel() string {
	return o.Label
}

func (o Object) In(x, y int) bool {
	return o.Image.RGBA64At(x-o.X, y-o.Y).A > 0
}

func (o Object) GetFont() font.Face {
	return o.Font
}
