package system

import "github.com/Lolodin/infclient/internal/entity"

type System struct {
	Entities   []*entity.Entity
	Components []string
}

func (s *System) AddEntity(e *entity.Entity) {
	s.Entities = append(s.Entities, e)
}

func (s System) GetComponents() []string {
	return s.Components
}