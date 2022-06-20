package entity

import (
	"github.com/Lolodin/infclient/internal/component"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
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
	e.Interactive = component.NewInteractiveComponent(&component.Object{text.BoundString(options.Font, options.Text)})

	// No text
	if options.Text == "" {
		return e, nil
	}

	e.Text = &component.TextComponent{
		Content: options.Text,
		Color:   options.Color,
		Font:    options.Font,
	}

	return e, nil
}

func getTextRect(f font.Face, content string) image.Rectangle {
	return text.BoundString(f, content)
}
