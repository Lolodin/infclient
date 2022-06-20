package entity

import (
	"github.com/Lolodin/infclient/internal/component"
)

func NewCursorEntity() (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{X: 10, Y: 10}
	e.Size = &component.SizeComponent{Width: 0.1, Height: 0.01}

	return e, nil
}
