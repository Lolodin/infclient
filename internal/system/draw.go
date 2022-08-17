package system

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/component"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"sort"
)

type SystemDrawAdapter struct {
	Entities   map[int][]*entity.Entity
	Components []string
}

func (s *SystemDrawAdapter) Update(e *component.PositionComponent, z int) {
	for i, e2 := range s.Entities[e.Z] {
		if e2.Position == e {
			s.Entities[e.Z] = append(s.Entities[e.Z][:i], s.Entities[e.Z][i+1:]...)
			s.Entities[z] = append(s.Entities[z], e2)
		}
	}

}
func (s *SystemDrawAdapter) AddEntity(e *entity.Entity) {
	s.Entities[e.Position.Z] = append(s.Entities[e.Position.Z], e)
}

func (s SystemDrawAdapter) GetComponents() []string {
	return s.Components
}

func (s *SystemDrawAdapter) Remove(e *entity.Entity) bool {
	for i, entity := range s.Entities[e.Position.Z] {
		if entity == e {
			s.Entities[e.Position.Z] = append(s.Entities[e.Position.Z][:i], s.Entities[e.Position.Z][i+1:]...)
			return true
		}
	}
	return false
}

// Система отрисовки игровых объектов и тестов
type drawSystem struct {
	LocalImage *ebiten.Image
	*SystemDrawAdapter
}

func NewDrawSystem(s *kernel.State, syst *SystemDrawAdapter) *drawSystem {
	sys := &drawSystem{LocalImage: ebiten.NewImage(s.RenderWidth, s.RenderHeight), SystemDrawAdapter: syst}
	sys.Components = []string{
		"View",
		"Position",
		"Size",
	}
	return sys
}

func (sys *drawSystem) Load(s *kernel.State) {}

func (sys *drawSystem) Update(s *kernel.State) {}

func (sys *drawSystem) Draw(s *kernel.State, screen *ebiten.Image) {
	sys.LocalImage = ebiten.NewImageWithOptions(image.Rect(int(s.Camera.Position[0]), int(s.Camera.Position[1]), s.RenderWidth, s.RenderHeight), nil)
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})
	keys := sys.getEntityZLayouts()
	for _, key := range keys {
		for _, e := range sys.Entities[key] {
			// Draw image
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(e.Position.X, e.Position.Y)

			if e.View.IsDraw {
				sys.LocalImage.DrawImage(
					e.View.Image.SubImage(*e.View.Frames[e.View.Frame]).(*ebiten.Image),
					options,
				)

			}

			if e.Text == nil {
				s.Camera.Render(sys.LocalImage, screen)
				continue
			}
			// No text? nothing to do
			if e.Text.TextField != nil {
				e.Text.TextField.Draw(sys.LocalImage)
			}
			if e.Text.InputField != nil {
				if e.Text.InputField.IsActive {
					e.Text.InputField.Draw(sys.LocalImage)
				}

			}

			s.Camera.Render(sys.LocalImage, screen)
			//	ebitenutil.DrawRect(sys.LocalImage, s.Camera.Position[0], s.Camera.Position[1], float64(s.RenderWidth), float64(s.RenderHeight), color.NRGBA{0x00, 0x40, 0x80, 0xff})
		}
	}

	worldX, worldY := s.Camera.ScreenToWorld(int(s.CursorX), int(s.CursorY))
	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f\nCursor X: %f\nCursor Y:%f\nCamera X: %f\nCamera Y:%f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		worldX,
		worldY,
		s.Camera.Position[0],
		s.Camera.Position[1],
	))

}

func (sys *drawSystem) Enter(s *kernel.State) {

}

func (sys *drawSystem) Exit(s *kernel.State) {}

//Отрисовка в глубину, по Z оси
func (sys *drawSystem) getEntityZLayouts() []int {
	i := 0
	keys := make([]int, len(sys.Entities))
	for k := range sys.Entities {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}
