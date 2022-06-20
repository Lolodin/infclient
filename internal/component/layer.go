package component

import "github.com/hajimehoshi/ebiten/v2"

type LayerComponent struct {
	Image *ebiten.Image
	Z     int
}

func NewLayerComponent(z int) *LayerComponent {
	return &LayerComponent{Image: &ebiten.Image{}, Z: z}
}
