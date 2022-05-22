package entity

import "github.com/Lolodin/infclient/internal/component"

type Entity struct {
	Position *component.PositionComponent
	Text     *component.TextComponent
	Input    *component.InputComponent
	Size     *component.SizeComponent
}

func NewEntity() *Entity {
	return &Entity{}
}
