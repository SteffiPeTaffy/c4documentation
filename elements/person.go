package elements

import (
	"fmt"
)

type Person struct {
	*C4Element
	description string
	external    bool
}

func NewPerson(name string) *Person {
	person := Person{
		C4Element:   &C4Element{
			C4BaseElement:   &C4BaseElement{
				Name:              name,
				OutgoingRelations: []*C4Relation{},
			},
		},
	}
	person.C4Writer = func() string {
		return person.toC4PlantUMLString()
	}
	return &person
}

func (p *Person) Description(description string) *Person {
	p.description = description
	return p
}

func (p *Person) External(external bool) *Person {
	p.external = external
	return p
}

func (p *Person) BelongsTo(parent *SystemBoundary) *Person {
	p.C4Element.BelongsTo(parent)
	return p
}

func (p *Person) RelatesTo(to ElementWithBase, label string, technology string) *Person {
	p.C4Element.RelatesTo(to,label,technology)
	return p
}

func (p *Person) toC4PlantUMLString() string {
	if p.external {
		return fmt.Sprintf("Person_Ext(%v, '%s', '%s')\n", p.Alias(), p.Name, p.description)
	}
	return fmt.Sprintf("Person(%v, '%s', '%s')\n", p.Alias(), p.Name, p.description)
}
