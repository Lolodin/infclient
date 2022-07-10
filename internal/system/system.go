package system

import (
	"github.com/Lolodin/infclient/internal/entity"
	"reflect"
)

type System struct {
	Entities   map[int][]*entity.Entity
	Components []string
}

// z слой в котором отрисовывается
func (s *System) AddEntity(e *entity.Entity, z int) {
	s.Entities[z] = append(s.Entities[z], e)
}

func (s *System) DeleteEntity(e *entity.Entity) {
	for k, entities := range s.Entities {
		for i, e2 := range entities {
			if e2 == e {
				s.Entities[k] = append(s.Entities[k][:i], s.Entities[k][i+1:]...)
			}

		}
	}
}

func (s *System) AddEntityInRealTime(e *entity.Entity, z int) {
	hasComponents := false
	entityReflection := reflect.Indirect(reflect.ValueOf(e))
	components := s.GetComponents()

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
		s.AddEntity(e, z)
	}
}

func (s System) GetComponents() []string {
	return s.Components
}
