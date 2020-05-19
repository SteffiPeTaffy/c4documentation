package elements

import (
	"fmt"
)

type System struct {
	NamedElement
	Description string
	External    bool
}

func NewSystem(name string, description string, external bool) *System {
	return &System{
		NamedElement: NamedElement{Name: name},
		Description:  description,
		External:     external,
	}
}

func (s *System) ToPlantUMLString() string {
	if s.External {
		return fmt.Sprintf("System_Ext(%v, '%s', '%s')\n", s.Alias(), s.Name, s.Description)
	}
	return fmt.Sprintf("System(%v, '%s', '%s')\n", s.Alias(), s.Name, s.Description)
}
