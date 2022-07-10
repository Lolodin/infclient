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

//Для отрисовки диалоговых окон, кнопок меню и полей для ввода
//TODO возможно стоит разделить на 2 отдельных компонента
type TextComponent struct {
	Content string
	Color   color.NRGBA
	Font    font.Face

	InputField *Input
	TextField  *messeji.TextField
}

type Input struct {
	IsActive bool
	*messeji.InputField
}

func NewInput(face font.Face) *Input {
	i := &Input{}

	i.InputField = messeji.NewInputField(face)
	i.InputField.SetAutoHideScrollBar(false)
	i.InputField.SetScrollBarVisible(false)
	return i
}
