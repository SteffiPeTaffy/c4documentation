package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
	"strings"
)

type C4ContextDiagram struct {
	name  string
	model *elements.C4Model
}

func NewContextDiagram(name string, model *elements.C4Model) *C4ContextDiagram {
	return &C4ContextDiagram{
		name:  name,
		model: model,
	}
}

func (c *C4ContextDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!include https://raw.githubusercontent.com/adrianvlupu/C4-PlantUML/latest/C4_Context.puml\n")
	b.WriteString("LAYOUT_TOP_DOWN()\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	all := func(element *elements.C4Element) bool {
		return true
	}
	boundaryView := c.model.CreateBoundaryView(all)

	for _, child := range boundaryView.Children {
		b.WriteString(child.C4Writer())
	}

	for _, topLevelBoundary := range boundaryView.NestedBoundaries {
		b.WriteString(fmt.Sprintf("System(%s, %s, %s)\n", topLevelBoundary.ElementInView.Alias(), topLevelBoundary.ElementInView.Name, ""))
	}

	systemRelations := findSystemRelations(c.model.Elements)
	for _, relation := range systemRelations {
		b.WriteString(relation.ToC4PlantUMLString())
	}

	b.WriteString("@enduml")

	return b.String()
}

func findSystemRelations(elems []*elements.C4Element) []*elements.C4Relation {
	relationsMap := make(map[elements.C4Alias][]*elements.C4Relation)
	for _, element := range elems {
		for _, relation := range element.OutgoingRelations {
			systemRelation := &elements.C4Relation {
				From:       findRoot(relation.From),
				To:         findRoot(relation.To),
				Label:      relation.Label,
				Technology: relation.Technology,
			}
			relationsMap = addToSystemRelations(relationsMap, systemRelation)
		}
	}

	allRelations := make([]*elements.C4Relation, 0, len(relationsMap))
	for  _, value := range relationsMap {
		allRelations = append(allRelations, value...)
	}

	return allRelations
}

func addToSystemRelations(relations map[elements.C4Alias][]*elements.C4Relation, newRelation *elements.C4Relation) map[elements.C4Alias][]*elements.C4Relation {
	if relations[newRelation.From.Alias()] == nil {
		relations[newRelation.From.Alias()] = []*elements.C4Relation{newRelation}
		return relations
	}

	for _, toRelation := range relations[newRelation.From.Alias()] {
		if toRelation.To.Alias() == newRelation.To.Alias() {
			toRelation.Label = concatUniques(toRelation.Label, newRelation.Label)
			toRelation.Technology = concatUniques(toRelation.Technology, newRelation.Technology)
			return relations
		}
	}

	relations[newRelation.From.Alias()] = append(relations[newRelation.From.Alias()], newRelation)
	return relations
}

func concatUniques(label string, newLabel string) string {
	existingLabels := strings.Split(label, ", ")
	existingLabels = unique(append(existingLabels, newLabel))

	return strings.Join(existingLabels, ", ")
}

func unique(list []string) []string {
	keys := make(map[string]bool)
	var tmp []string
	for _, entry := range list {
		if _, value := keys[entry]; !value  && entry != ""{
			keys[entry] = true
			tmp = append(tmp, entry)
		}
	}
	return tmp
}

func findRoot(element *elements.C4BaseElement) (root *elements.C4BaseElement) {
	root = element
	for root.Parent != nil {
		root = root.Parent.C4BaseElement
	}
	return
}
