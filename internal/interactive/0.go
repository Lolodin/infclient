package interactive

import (
	"github.com/Lolodin/infclient/internal/entity"
	"github.com/Lolodin/infclient/internal/kernel"
	"github.com/hajimehoshi/ebiten/v2"
	"reflect"
)

type world struct {
	name     string
	systems  []kernel.GameSystem
	entities []*entity.Entity
}

func (w *world) Update(s *kernel.State) {
	for _, sys := range w.systems {
		sys.Update(s)
	}
}

func (w *world) Draw(s *kernel.State, screen *ebiten.Image) {
	for _, sys := range w.systems {
		sys.Draw(s, screen)
	}
}

func (w *world) Enter(s *kernel.State) {
	for _, sys := range w.systems {
		sys.Enter(s)
	}
}

func (w *world) Exit(s *kernel.State) {
	for _, sys := range w.systems {
		sys.Enter(s)
	}
}

func (w *world) Load(s *kernel.State) {
	for _, sys := range w.systems {
		sys.Load(s)
	}
}

func (w *world) Name() string {
	return w.name
}

// Add entities to systems based on their components. This is an
// expensive function and should be used sparingly. Ideally after
// multiple system and entity updates.
func (w *world) updateSystems() {
	// Assume entity to be suitable by default
	hasComponents := true

	for _, system := range w.systems {
		for _, entity := range w.entities {
			entityReflection := reflect.Indirect(reflect.ValueOf(entity))
			components := system.GetComponents()

			// A system w/o components should have no entities
			if len(components) == 0 {
				hasComponents = false
			}

			// Check if entity has required components
			for _, component := range components {
				field := entityReflection.FieldByName(component)
				if field.IsNil() {
					hasComponents = false
				}
			}

			// The entity is suitable and we can add it
			if hasComponents {
				system.AddEntity(entity, entity.Layer.Z)
			}

			// Reset for next iteration
			hasComponents = true
		}
	}
}
