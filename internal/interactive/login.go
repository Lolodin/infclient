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
	observ := &system.SystemDrawAdapter{Entities: map[int][]*entity.Entity{}}
	w.systems = []kernel.GameSystem{
		system.NewDrawSystem(s, observ),
		system.NewInputSystem(s),
		system.NewCursorSystem(s),
		system.NewInteractiveSystem(s),
	}

	cursor, err := entity.NewCursorEntity()
	if err != nil {
		panic(fmt.Sprintf("creating cursor entity: %s", err))
	}
	background, err := entity.NewImageEntity("./internal/resource/scroll.png", observ)
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
	loginField.Position.Observer = observ
	loginField.Position.Z = 2

	loginField.Text.InputField.SetHandleKeyboard(true)
	loginField.Text.InputField.SetSelectedFunc(func() (accept bool) {
		if loginField.Text.InputField.IsActive {
			fmt.Println(loginField.Text.InputField.Text())
			return true
		}
		return false
	})

	buffer, err := entity.NewButtonWithTextFieldEntity(&entity.ButtonEntityOptions{
		X:          (float64(s.RenderWidth) / 2.1) - (buttonWidth / 2),
		Y:          buttonYStart - 30,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    0,
		Text:       s.Lang.TransWithOut("enter_login"),
		Font:       s.Fonts["std"],
		Color:      color.NRGBA{uint8(A), uint8(B), uint8(C), uint8(D)},
		IsCentered: true,
	})
	buffer.Position.Z = 2
	buffer.Position.Observer = observ

	buffer2, err := entity.NewButtonWithTextFieldEntity(&entity.ButtonEntityOptions{
		X:          (float64(s.RenderWidth) / 2.1) - (buttonWidth / 2),
		Y:          buttonYStart - 30,
		Width:      buttonWidth,
		Height:     buttonHeight,
		Padding:    0,
		Text:       s.Lang.TransWithOut("enter_pass"),
		Font:       s.Fonts["std"],
		Color:      color.NRGBA{uint8(A), uint8(B), uint8(C), uint8(D)},
		IsCentered: true,
	})
	buffer2.Position.Observer = observ

	loginField.Text.InputField.SetSelectedFunc(func() (accept bool) {
		if loginField.Text.InputField.IsActive {
			s.Auth.SetLogin(loginField.Text.InputField.Text())
			for _, gameSystem := range w.systems {
				if ok := gameSystem.Remove(loginField); ok {
					fmt.Println("delete")
				}
			}
			for _, gameSystem := range w.systems {
				if ok := gameSystem.Remove(buffer); ok {
					fmt.Println("delete")
				}
			}

			buffer2.Position.Update(2) //Сделать тригер который перебрасывает в нужный слой
			return true
		}

		return false
	})
	loginField.Interactive.ClickEvent = func() {
		loginField.Text.InputField.IsActive = !loginField.Text.InputField.IsActive
	}

	s.MouseInputs = map[ebiten.MouseButton]component.Control{ebiten.MouseButtonLeft: component.ControlLeftClick}
	w.entities = []*entity.Entity{
		background,
		cursor,
		loginField,
		buffer,
		buffer2,
	}

	w.updateSystems()

	return w
}
