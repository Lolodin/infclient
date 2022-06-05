package system

import (
	"fmt"
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"sort"
)

type drawSystem struct {
	System
}

func NewDrawSystem(s *kernel.State) *drawSystem {
	sys := &drawSystem{}
	sys.System.Entities = map[int][]*entity.Entity{}
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
	screen.Fill(color.NRGBA{0x00, 0x40, 0x80, 0xff})

	for _, z := range sys.getEntityZLayouts() {
		for _, e := range sys.Entities[z] {
			// Draw image
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(e.Position.X, e.Position.Y)

			if e.View.IsDraw {
				screen.DrawImage(
					e.View.Image.SubImage(*e.View.Frames[e.View.Frame]).(*ebiten.Image),
					options,
				)
			}

			// No text? nothing to do
			if e.Text == nil {
				continue
			}

			// Draw lines of text
			for i, line := range e.Text.Lines {
				x := line.X + e.Text.Padding
				y := (e.Size.Height / 2) - float64((e.Text.LineHeight*len(e.Text.Lines))/2) + float64((i+1)*e.Text.LineHeight)

				text.Draw(screen, line.Content, e.Text.Font, int(e.Position.X+x), int(e.Position.Y+y), e.Text.Color)
			}
		}

		ebitenutil.DebugPrint(screen, fmt.Sprintf(
			"TPS: %0.2f\nFPS: %0.2f\nCursor X: %f\nCursor Y:%f",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			s.CursorX,
			s.CursorY,
		))
	}

}

func (sys *drawSystem) Enter(s *kernel.State) {}

func (sys *drawSystem) Exit(s *kernel.State) {}

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
