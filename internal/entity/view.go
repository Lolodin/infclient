package entity

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/component"
)

func NewImageEntity(path string) (*Entity, error) {
	e := NewEntity()
	e.Layer = component.NewLayerComponent(0)
	e.Position = &component.PositionComponent{X: 0, Y: 0}
	e.Size = &component.SizeComponent{}
	view, err := component.NewAppearanceComponent(path, "")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.View = view
	return e, nil
}
