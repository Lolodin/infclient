package entity

import (
	"github.com/Lolodin/infclient/internal/component"
	"image"
)

type Entity struct {
	Position    *component.PositionComponent
	Text        *component.TextComponent
	Size        *component.SizeComponent
	View        *component.ViewComponent
	Interactive *component.InteractiveComponent
}

func NewEntity() *Entity {
	return &Entity{}
}

func (e Entity) IsClick(data *component.InputData) bool {
	return image.Rect(data.X, data.Y, data.X+1, data.Y+1).In(e.GetRec())
}

func (e Entity) GetRec() image.Rectangle {
	return image.Rect(int(e.Position.X), int(e.Position.Y), int(e.Position.X+e.Size.Width), int(e.Position.Y+e.Size.Height))
}
