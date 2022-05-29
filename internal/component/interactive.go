package component

type InteractiveComponent struct {
}

func NewInteractiveComponent() *InteractiveComponent {
	return &InteractiveComponent{}
}

func (r *InteractiveComponent) IsClick(data *InputData, pos *PositionComponent, size *SizeComponent) bool {
	return float64(data.X) >= pos.X && float64(data.X) <= size.Width+pos.X &&
		float64(data.Y) >= pos.Y && float64(data.Y) <= pos.Y+size.Height
}
