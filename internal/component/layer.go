package component

type LayerComponent struct {
	Z int
}

func NewLayerComponent(z int) *LayerComponent {
	return &LayerComponent{z}
}
