package system

import (
	"github.com/Lolodin/infclient/internal/component"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type inputSystem struct {
	System
}

func NewInputSystem(s *kernel.State) *inputSystem {
	sys := &inputSystem{}
	sys.System.Entities = map[int][]*entity.Entity{}
	return sys
}

func (sys *inputSystem) Load(s *kernel.State) {}

func (sys *inputSystem) Update(s *kernel.State) {
	for mouseInput, control := range s.MouseInputs {
		if inpututil.IsMouseButtonJustReleased(mouseInput) {
			ix, iy := ebiten.CursorPosition()
			s.SetControl(control, &component.InputData{ix, iy})
		}
	}

}

func (sys *inputSystem) Draw(s *kernel.State, screen *ebiten.Image) {}

func (sys *inputSystem) Enter(s *kernel.State) {}

func (sys *inputSystem) Exit(s *kernel.State) {}
