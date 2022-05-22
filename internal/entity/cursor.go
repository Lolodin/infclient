package entity

import (
	"github.com/Lolodin/infclient/internal/component"
)

func NewCursorEntity() (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: 10, Y: 10}
	e.Size = &component.SizeComponent{Width: 10, Height: 10}

	return e, nil
}
