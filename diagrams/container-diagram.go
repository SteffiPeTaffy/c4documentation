package diagrams

import (
	"bytes"
	"c4documentation/elements"
	"fmt"
)

type ContainerDiagram struct {
	Name string
	Elements []elements.NamedElement
}

func (c *ContainerDiagram) Add(element elements.NamedElement) *ContainerDiagram {
	c.Elements = append(c.Elements, element)
	return c
}

func (c *ContainerDiagram) ToPlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.Name))
	b.WriteString("!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Container.puml\n")
	b.WriteString("LAYOUT_TOP_DOWN\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	for _, element := range c.Elements {
		b.WriteString(element.C4Writer(&element))
	}

	b.WriteString("@enduml")

	return b.String()
}