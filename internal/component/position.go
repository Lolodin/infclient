package component

type Observer interface {
	Update(component *PositionComponent, newZ int)
}

// PositionComponent Указывает позицию объекта в мире по X, Y.
// Z используется как слой  при отрисовке объекта

type PositionComponent struct {
	Observer
	X, Y float64
	Z    int // layer
}

func NewPositionComponent(x, y float64) *PositionComponent {
	return &PositionComponent{
		X: x,
		Y: y,
	}
}
func (p *PositionComponent) Update(Z int) {
	p.Observer.Update(p, Z)
	p.Z = Z
}
