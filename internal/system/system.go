package system

import (
	"github.com/Lolodin/infclient/internal/entity"
)

type System struct {
	Entities   []*entity.Entity
	Components []string
}

// z слой в котором отрисовывается
func (s *System) AddEntity(e *entity.Entity) {
	s.Entities = append(s.Entities, e)
}

func (s System) GetComponents() []string {
	return s.Components
}

func (s *System) Remove(e *entity.Entity) bool {
	for i, entity := range s.Entities {
		if entity == e {
			s.Entities = append(s.Entities[:i], s.Entities[i+1:]...)
			return true
		}
	}
	return false
}
