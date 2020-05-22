package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

type C4ContainerDiagram struct {
	elements.C4Printable
	name  string
	model elements.C4Model
}

func NewContainerDiagram(name string, model elements.C4Model) *C4ContainerDiagram {
	return &C4ContainerDiagram{
		name:  name,
		model: model,
	}
}

func (c *C4ContainerDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Container.puml\n")
	b.WriteString("LAYOUT_TOP_DOWN\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	for _, element := range c.model.Elements {
		b.WriteString(element.C4Writer())
	}

	for _, container := range c.model.Boundaries {
		b.WriteString(container.C4Writer())
	}

	for _, element := range c.model.Elements {
		for _, relation := range element.OutgoingRelations {
			b.WriteString(relation.ToC4PlantUMLString())
		}
	}

	for _, boundary := range c.model.Boundaries {
		boundary.VisitElements(func(element elements.C4NodeElement) (done bool) {
			for _, relation := range element.OutgoingRelations {
				b.WriteString(relation.ToC4PlantUMLString())
			}
			return false
		})
	}

	b.WriteString("@enduml")

	return b.String()
}
