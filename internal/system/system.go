package system

import (
	"github.com/Lolodin/infclient/internal/entity"
)

type System struct {
	Entities   map[int][]*entity.Entity
	Components []string
}

// z слой в котором отрисовывается
func (s *System) AddEntity(e *entity.Entity, z int) {
	s.Entities[z] = append(s.Entities[z], e)
}

func (s System) GetComponents() []string {
	return s.Components
}

func (s *System) Remove(e *entity.Entity) bool {
	for k, entities := range s.Entities {
		for i, e2 := range entities {
			if e2 == e {
				s.Entities[k] = append(s.Entities[k][:i], s.Entities[k][i+1:]...)
				return true
			}
		}
	}
	return false
}
