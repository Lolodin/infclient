package interactive

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/Lolodin/infclient/internal/system"
)

type titleWorld struct {
	world
}

func NewTitleWorld(s *kernel.State) *titleWorld {
	w := &titleWorld{}
	w.name = "title"

	w.systems = []kernel.GameSystem{
		system.NewDrawSystem(s),
		system.NewInputSystem(s),
		system.NewCursorSystem(s),
	}

	cursor, err := entity.NewCursorEntity()
	if err != nil {
		panic(fmt.Sprintf("creating cursor entity: %s", err))
	}

	buttonWidth := 74.0
	buttonHeight := 30.0
	buttonYStart := 120.0
	startButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:          (float64(s.RenderWidth) / 2) - (buttonWidth / 2),
		Y:          buttonYStart,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    10,
		Text:       s.Lang.TransWithOut("Login"),
		Font:       s.Fonts["Comic_Sans_MS.ttf"],
		IsCentered: true,
	})
	if err != nil {
		panic(fmt.Sprintf("creating title button entity: %s", err))
	}

	w.entities = []*entity.Entity{
		startButton,
		cursor,
	}
	w.updateSystems()

	return w
}
