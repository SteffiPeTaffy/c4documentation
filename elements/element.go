package elements

import (
	"regexp"
)

type C4Printable interface {
	ToC4PlantUMLString() string
}

type C4Alias string

type C4PlantUMLAlias interface {
	Alias() C4Alias
}

type C4BaseElement struct {
	Name string
	OutgoingRelations []*C4Relation
	Parent   *SystemBoundary
}

func (n *C4BaseElement) GetBase() *C4BaseElement {
	return n
}

func (n *C4BaseElement) Alias() C4Alias {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return C4Alias(reg.ReplaceAllString(n.Name, ""))
}

func (n *C4BaseElement) RelatesTo(to interface{}, label string, technology string) {
	n.OutgoingRelations = append(n.OutgoingRelations, &C4Relation{
		From:       n,
		To:         to.(*C4BaseElement),
		Label:      label,
		Technology: technology,
	})
}

func (n *C4BaseElement) BelongsTo(parent *SystemBoundary) {
	n.Parent = parent
}


type C4Element struct {
	*C4BaseElement
	C4Writer func() string
}

func (n *C4Element) WritePUML() string {
	return n.C4Writer()
}

func (n *C4Element) BelongsTo(parent *SystemBoundary) {
	n.Parent = parent
}

func (n *C4Element) RelatesTo(to ElementWithBase, label string, technology string) {
	n.OutgoingRelations = append(n.OutgoingRelations, &C4Relation{
		From:       n.C4BaseElement,
		To:         to.GetBase(),
		Label:      label,
		Technology: technology,
	})
}
