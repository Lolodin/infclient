package interactive

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/component"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/Lolodin/infclient/internal/system"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type titleWorld struct {
	world
}

func NewTitleWorld(s *kernel.State) *titleWorld {
	w := &titleWorld{}
	w.name = "main"

	w.systems = []kernel.GameSystem{
		system.NewDrawSystem(s),
		system.NewInputSystem(s),
		system.NewCursorSystem(s),
		system.NewInteractiveSystem(s),
	}

	cursor, err := entity.NewCursorEntity()
	if err != nil {
		panic(fmt.Sprintf("creating cursor entity: %s", err))
	}
	s.MouseInputs = map[ebiten.MouseButton]component.Control{ebiten.MouseButtonLeft: component.ControlLeftClick}

	buttonWidth := 200.0
	buttonHeight := 30.0
	buttonYStart := 120.0
	A, B, C, D := color.White.RGBA()
	startButton, err := entity.NewButtonEntity(&entity.ButtonEntityOptions{
		X:          (float64(s.RenderWidth) / 2) - (buttonWidth / 2),
		Y:          buttonYStart,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    0,
		Text:       s.Lang.TransWithOut("login"),
		Font:       s.Fonts["std"],
		Color:      color.NRGBA{uint8(A), uint8(B), uint8(C), uint8(D)},
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
