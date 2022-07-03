package component

import (
	"code.rocketnine.space/tslocum/messeji"
	"image/color"

	"golang.org/x/image/font"
)

type TextLine struct {
	Content string
	X       float64
}

type TextComponent struct {
	Content string
	Color   color.NRGBA
	Font    font.Face

	InputField *messeji.InputField
	TextField  *messeji.TextField
}
