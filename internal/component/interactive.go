package component

import "image"

type Interactor interface {
	IsClick(data *InputData, pos *PositionComponent, size *SizeComponent) bool
}

type InteractiveComponent struct {
	Interactor
}

func NewInteractiveComponent(interactor Interactor) *InteractiveComponent {
	return &InteractiveComponent{interactor}
}

type Button struct {
	image.Rectangle
}

func (r *Button) IsClick(data *InputData, pos *PositionComponent, size *SizeComponent) bool {
	return float64(data.X) >= pos.X && float64(data.X) <= size.Width+pos.X &&
		float64(data.Y) >= pos.Y-size.Height/3 && float64(data.Y) <= pos.Y+size.Height/3
}

type Object struct {
	image.Rectangle
}

func (r *Object) IsClick(data *InputData, pos *PositionComponent, size *SizeComponent) bool {
	return r.RGBA64At(data.X-int(pos.X), data.Y-int(pos.Y)).A > 0
}
