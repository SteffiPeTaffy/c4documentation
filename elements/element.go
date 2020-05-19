package elements

import (
	"regexp"
)

type Alias string

type AliasElement interface {
	Alias() Alias
}

type PlantUMLElement interface {
	ToPlantUMLString() string
}

type NamedElement struct {
	Name              string
	OutgoingRelations []Relation
	PlantUMLElement
	AliasElement
}

func (n *NamedElement) Alias() Alias {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return Alias(reg.ReplaceAllString(n.Name, ""))
}

func (n *NamedElement) RelatesTo(element AliasElement, label string, technology string) *NamedElement {
	n.OutgoingRelations = append(n.OutgoingRelations, Relation{
		From:       *n,
		To:         element,
		Label:      label,
		Technology: technology,
	})
	return n
}
