package elements

import (
	"fmt"
)

type System struct {
	*C4Element
	description string
	external    bool
}

func NewSystem(name string) *System {
	system := System{
		C4Element:   &C4Element{
			C4BaseElement:   &C4BaseElement{
				Name:              name,
				OutgoingRelations: []*C4Relation{},
			},
		},
	}
	system.C4Writer = func() string {
		return system.toC4PlantUMLString()
	}
	return &system
}

func (s *System) Description(description string) *System {
	s.description = description
	return s
}

func (s *System) External(external bool) *System {
	s.external = external
	return s
}


func (s *System) toC4PlantUMLString() string {
	if s.external {
		return fmt.Sprintf("System_Ext(%v, '%s', '%s')\n", s.Alias(), s.Name, s.description)
	}
	return fmt.Sprintf("System(%v, '%s', '%s')\n", s.Alias(), s.Name, s.description)
}
