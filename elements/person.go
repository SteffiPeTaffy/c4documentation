package elements

import (
	"fmt"
)

type Person struct {
	NamedElement
	Description string
	External    bool
}

func NewPerson(name string, description string, external bool) *Person {
	return &Person{
		NamedElement: NamedElement{Name: name},
		Description:  description,
		External:     external,
	}
}

func (p *Person) ToPlantUMLString() string {
	if p.External {
		return fmt.Sprintf("Person_Ext(%v, '%s', '%s')", p.Alias(), p.Name, p.Description)
	}
	return fmt.Sprintf("Person(%v, '%s', '%s')\n", p.Alias(), p.Name, p.Description)
}
