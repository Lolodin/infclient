package system

import (
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"

	"github.com/hajimehoshi/ebiten/v2"
)

// Система для работу с курсором
type cursorSystem struct {
	System
}

func NewCursorSystem(s *kernel.State) *cursorSystem {
	sys := &cursorSystem{}
	sys.System.Entities = []*entity.Entity{}

	return sys
}

func (sys *cursorSystem) Load(s *kernel.State) {}

func (sys *cursorSystem) Update(s *kernel.State) {
	for _, e := range sys.Entities {

		e.Position.X = s.CursorX
		e.Position.Y = s.CursorY

	}
}

func (sys *cursorSystem) Draw(s *kernel.State, screen *ebiten.Image) {}

func (sys *cursorSystem) Enter(s *kernel.State) {

}

func (sys *cursorSystem) Exit(s *kernel.State) {

}
