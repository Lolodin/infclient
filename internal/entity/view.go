package entity

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/component"
)

func NewImageEntity(path string, observer component.Observer) (*Entity, error) {
	e := NewEntity()
	e.Position = &component.PositionComponent{Observer: observer}
	e.Size = &component.SizeComponent{}
	view, err := component.NewViewComponent(path, "")
	if err != nil {
		return e, fmt.Errorf("creating appearance component: %s", err)
	}
	e.View = view
	return e, nil
}
