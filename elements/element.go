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

func (n *C4BaseElement) Alias() C4Alias {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return C4Alias(reg.ReplaceAllString(n.Name, ""))
}

func (n *C4BaseElement) RelatesTo(to interface{}, label string, technology string) *C4BaseElement {
	n.OutgoingRelations = append(n.OutgoingRelations, &C4Relation{
		From:       n,
		To:         to.(*C4BaseElement),
		Label:      label,
		Technology: technology,
	})
	return n
}

func (n *C4BaseElement) BelongsTo(parent *SystemBoundary) *C4BaseElement {
	n.Parent = parent
	return n
}

func (n *C4BaseElement) Build() *C4BaseElement {
	return n
}

type C4Element struct {
	*C4BaseElement
	C4Writer func() string
}

func (n *C4Element) BelongsTo(parent *SystemBoundary) *C4Element {
	n.Parent = parent
	return n
}

func (n *C4Element) RelatesTo(to interface{}, label string, technology string) *C4Element {
	var toElement *C4BaseElement

	if val, ok := to.(*C4BaseElement); ok {
		toElement = val
	} else if val, ok := to.(*C4Element); ok{
		toElement = val.C4BaseElement
	} else if val, ok := to.(*SystemBoundary); ok {
		toElement = val.C4BaseElement
	} else {
		return n
	}

	n.OutgoingRelations = append(n.OutgoingRelations, &C4Relation{
		From:       n.C4BaseElement,
		To:         toElement,
		Label:      label,
		Technology: technology,
	})
	return n
}

func (n *C4Element) Build() *C4Element {
	return n
}

