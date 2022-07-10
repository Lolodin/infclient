package system

import (
	"github.com/Lolodin/infclient/internal/component"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/hajimehoshi/ebiten/v2"
)

type interactiveSystem struct {
	System
}

func NewInteractiveSystem(s *kernel.State) *interactiveSystem {
	sys := &interactiveSystem{}
	sys.System.Entities = map[int][]*entity.Entity{}
	sys.Components = []string{
		"Interactive",
	}
	return sys
}

func (sys *interactiveSystem) Load(s *kernel.State) {}

func (sys *interactiveSystem) Update(s *kernel.State) {
	for _, sl := range sys.Entities {
		for _, e := range sl {
			if data, ok := s.Controls[component.ControlLeftClick]; ok {
				if e.IsClick(data) {
					if e.Interactive != nil {
						e.Interactive.ClickEvent()
					}

				}
			}
			if e.Text == nil {
				continue
			}
			if e.Text.TextField != nil {
				e.Text.TextField.SetRect(e.GetRec())
			}
			if e.Text.InputField != nil {
				e.Text.InputField.SetRect(e.GetRec())
				if e.Text.InputField.IsActive {
					e.Text.InputField.Update()
				}
			}

		}
		s.Controls = map[component.Control]*component.InputData{}
	}
}

func (sys *interactiveSystem) Draw(s *kernel.State, screen *ebiten.Image) {}

func (sys *interactiveSystem) Enter(s *kernel.State) {}

func (sys *interactiveSystem) Exit(s *kernel.State) {}
