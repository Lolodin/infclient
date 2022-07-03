package component

// PositionComponent Указывает позицию объекта в мире по X, Y.
// Z используется как слой  при отрисовке объекта
type PositionComponent struct {
	X, Y float64
	Z    int // layer
}

func NewPositionComponent(x, y float64) *PositionComponent {
	return &PositionComponent{
		X: x,
		Y: y,
	}
}
