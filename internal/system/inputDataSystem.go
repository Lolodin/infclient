package system

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/hajimehoshi/ebiten/v2"
)

type inputDataSystem struct {
	System
}

func NewInputDataSystem(s *kernel.State) *interactiveSystem {
	sys := &interactiveSystem{}
	sys.System.Entities = map[int][]*entity.Entity{}
	sys.Components = []string{
		"Text",
	}
	return sys
}

func (sys *inputDataSystem) Load(s *kernel.State) {}

func (sys *inputDataSystem) Update(s *kernel.State) {
	for _, sl := range sys.Entities {
		for _, e := range sl {
			fmt.Println(e)
		}
	}
}

func (sys *inputDataSystem) Draw(s *kernel.State, screen *ebiten.Image) {}

func (sys *inputDataSystem) Enter(s *kernel.State) {}

func (sys *inputDataSystem) Exit(s *kernel.State) {}
