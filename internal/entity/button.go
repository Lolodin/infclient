package entity

import (
	"code.rocketnine.space/tslocum/messeji"
	"github.com/Lolodin/infclient/internal/component"
	"golang.org/x/image/font"
	"image/color"
)

type ButtonEntityOptions struct {
	X, Y, Width, Height, Padding float64
	IsCentered                   bool
	Text                         string
	Color                        color.NRGBA
	Font                         font.Face
	Z                            int
}

func NewButtonEntity(options *ButtonEntityOptions) (*Entity, error) {
	e := NewEntity()

	e.Layer = component.NewLayerComponent(options.Z)
	e.Position = &component.PositionComponent{X: options.X, Y: options.Y}
	e.Size = &component.SizeComponent{Width: options.Width, Height: options.Height}
	e.View = &component.ViewComponent{}

	// No text
	if options.Text == "" {
		return e, nil
	}

	e.Text = &component.TextComponent{
		Content: options.Text,
		Color:   options.Color,
		Font:    options.Font,
	}
	e.Text.TextField = messeji.NewTextField(options.Font)
	e.Text.TextField.SetBackgroundColor(color.Alpha{A: 0})
	e.Text.TextField.SetForegroundColor(options.Color)
	e.Text.TextField.SetScrollBarVisible(false)
	e.Text.TextField.SetText(options.Text)

	return e, nil
}
