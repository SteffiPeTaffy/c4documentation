package elements

import (
	"fmt"
	"regexp"
)

type C4Printable interface {
	ToC4PlantUMLString() string
}

type C4PlantUMLAlias interface {
	Alias() C4Alias
}

type C4Node interface {
	RelatesTo(to C4NodeElement, label string, technology string) *C4NodeElement
}

type C4Boundary interface {
	Add(element C4NodeElement) *C4Boundary
	VisitElements(callback func(element C4NodeElement) (done bool))
}

type C4Alias string

type C4NodeElement struct {
	C4PlantUMLAlias
	C4Node
	Name              string
	OutgoingRelations []C4Relation
	C4Writer          func() string
}

func (n *C4NodeElement) Build() *C4NodeElement {
	return n
}

func (n *C4NodeElement) Alias() C4Alias {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return C4Alias(reg.ReplaceAllString(n.Name, ""))
}

func (n *C4NodeElement) RelatesTo(to C4PlantUMLAlias, label string, technology string) *C4NodeElement {
	n.OutgoingRelations = append(n.OutgoingRelations, C4Relation{
		From:       *n,
		To:         to,
		Label:      label,
		Technology: technology,
	})
	return n
}

type C4BoundaryElement struct {
	C4Boundary
	C4NodeElement
	elements   []C4NodeElement
	containers []C4BoundaryElement
}

func (c *C4BoundaryElement) VisitElements(callback func(element C4NodeElement) (done bool)) {
	for _, elem := range c.elements {
		done := callback(elem)
		if done {
			return
		}
	}

	for _, container := range c.containers {
		container.VisitElements(callback)
	}
}

func (n *C4BoundaryElement) Build() *C4BoundaryElement {
	return n
}

type Step struct {
	C4Printable
	From       C4NodeElement
	To         C4NodeElement
	Label      string
	Technology string
}

func (s *Step) ToC4PlantUMLString() string {
	return fmt.Sprintf("Rel(%v, '%v', '%s', '%s')\n", s.From.Alias(), s.To.Alias(), s.Label, s.Technology)
}

