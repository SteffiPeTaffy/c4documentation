package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

type C4ContainerDiagram struct {
	name  string
	model *elements.C4Model
}

func NewContainerDiagram(name string, model *elements.C4Model) *C4ContainerDiagram {
	return &C4ContainerDiagram{
		name:  name,
		model: model,
	}
}

func (c *C4ContainerDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!include https://raw.githubusercontent.com/adrianvlupu/C4-PlantUML/latest/C4_Container.puml\n\n")
	b.WriteString("LAYOUT_TOP_DOWN()\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	all := func(element *elements.C4Element) bool {
		return true
	}
	b.WriteString(c.model.CreateBoundaryView(all).ToC4PlantUMLString())

	for _, element := range c.model.Elements {
		for _, relation := range element.OutgoingRelations {
			b.WriteString(relation.ToC4PlantUMLString())
		}
	}

	b.WriteString("@enduml")

	return b.String()
}
