package component

import "fmt"

// InteractiveComponent Помечает объект как интерактивный, для его обработки в interactiveSystem
// Используется если нам нужно обрабатывать клик\наводку\коллизию и т.д.

type InteractiveComponent struct {
	ClickEvent
}
type ClickEvent func()

func NewInteractiveComponent() *InteractiveComponent {
	return &InteractiveComponent{func() {
		fmt.Println('x')
	}}
}
