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

type loginWorld struct {
	world
}

func NewLoginWorld(s *kernel.State) *loginWorld {
	w := &loginWorld{}
	w.name = "login"
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
	background, err := entity.NewImageEntity("./internal/resource/scroll.png")
	background.Position.Z = 1
	background.Position.X += 100

	if err != nil {
		panic(fmt.Sprintf("creating background entity: %s", err))
	}
	buttonWidth := 200.0
	buttonHeight := 30.0
	buttonYStart := 120.0
	A, B, C, D := color.Black.RGBA()
	loginField, err := entity.NewButtonWithTextInputFieldEntity(&entity.ButtonEntityOptions{
		X:          (float64(s.RenderWidth) / 2.1) - (buttonWidth / 2),
		Y:          buttonYStart,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    0,
		Text:       s.Lang.TransWithOut("login"),
		Font:       s.Fonts["std"],
		Color:      color.NRGBA{uint8(A), uint8(B), uint8(C), uint8(D)},
		IsCentered: true,
	})
	loginField.Position.Z = 2

	loginField.Text.InputField.SetHandleKeyboard(true)
	loginField.Text.InputField.SetSelectedFunc(func() (accept bool) {
		fmt.Println(loginField.Text.InputField.Text())
		return true
	})

	buffer, err := entity.NewButtonWithTextFieldEntity(&entity.ButtonEntityOptions{
		X:          (float64(s.RenderWidth) / 2.1) - (buttonWidth / 2),
		Y:          buttonYStart - 30,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    0,
		Text:       s.Lang.TransWithOut("login"),
		Font:       s.Fonts["std"],
		Color:      color.NRGBA{uint8(A), uint8(B), uint8(C), uint8(D)},
		IsCentered: true,
	})
	buffer.Position.Z = 2

	loginField.Text.InputField.SetSelectedFunc(func() (accept bool) {
		buffer.Text.TextField.SetText(loginField.Text.InputField.Text())
		return true
	})

	s.MouseInputs = map[ebiten.MouseButton]component.Control{ebiten.MouseButtonLeft: component.ControlLeftClick}
	w.entities = []*entity.Entity{
		background,
		cursor,
		loginField,
		buffer,
	}

	w.updateSystems()

	return w
}