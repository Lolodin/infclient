package entity

import "github.com/Lolodin/infclient/internal/component"

type Entity struct {
	Position    *component.PositionComponent
	Text        *component.TextComponent
	Size        *component.SizeComponent
	Interactive *component.InteractiveComponent
	View        *component.ViewComponent
	Layer       *component.LayerComponent
}

func NewEntity() *Entity {
	return &Entity{}
}

func (e Entity) IsClick(data *component.InputData) bool {
	return e.Interactive.IsClick(data, e.Position, e.Size)
}
