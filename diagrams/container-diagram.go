package diagrams

import (
	"bytes"
	"c4documentation/elements"
	"fmt"
)

type C4ContainerDiagram struct {
	elements.C4Printable
	name       string
	elements   []elements.C4NodeElement
	containers []elements.C4ContainerElement
}

func NewContainerDiagram (name string) *C4ContainerDiagram {
	return &C4ContainerDiagram {
		name:        name,
	}
}

func (c *C4ContainerDiagram) Add(element elements.C4NodeElement) *C4ContainerDiagram {
	c.elements = append(c.elements, element)
	return c
}

func (c *C4ContainerDiagram) AddSystemBoundary(container elements.C4ContainerElement) *C4ContainerDiagram {
	c.containers = append(c.containers, container)
	return c
}

func (c *C4ContainerDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Container.puml\n")
	b.WriteString("LAYOUT_TOP_DOWN\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	for _, element := range c.elements {
		b.WriteString(element.C4Writer())
	}

	for _, container := range c.containers {
		b.WriteString(container.C4Writer())
	}

	for _, element := range c.elements {
		for _, relation := range element.OutgoingRelations {
			b.WriteString(relation.ToC4PlantUMLString())
		}
	}

	for _, container := range c.containers {
		container.VisitElements(func(element elements.C4NodeElement) (done bool) {
			for _, relation := range element.OutgoingRelations {
				b.WriteString(relation.ToC4PlantUMLString())
			}
			return false
		})
	}

	b.WriteString("@enduml")

	return b.String()
}
