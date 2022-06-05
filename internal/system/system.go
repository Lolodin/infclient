package system

import "github.com/Lolodin/infclient/internal/entity"

type System struct {
	Entities   map[int][]*entity.Entity
	Components []string
}

// e сущность, z слой в котором отрисовывается
func (s *System) AddEntity(e *entity.Entity, z int) {
	s.Entities[z] = append(s.Entities[z], e)
}

func (s System) GetComponents() []string {
	return s.Components
}
