package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

type C4SequenceDiagram struct {
	elements.C4Printable
	name     string
	elements map[elements.C4Alias]elements.C4NodeElement
	sequence []elements.Step
}

func NewSequenceDiagram(name string) *C4SequenceDiagram {
	return &C4SequenceDiagram{
		name: name,
		elements: map[elements.C4Alias]elements.C4NodeElement{},
		sequence: []elements.Step{},
	}
}
func (c *C4SequenceDiagram) Next(from elements.C4NodeElement, to elements.C4NodeElement, label string, technology string) *C4SequenceDiagram {
	c.elements[from.Alias()] = from
	c.elements[to.Alias()] = to
	c.sequence = append(c.sequence, elements.Step{
		From:       from,
		To:         to,
		Label:      label,
		Technology: technology,
	})
	return c
}

func (c *C4SequenceDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!include https://raw.githubusercontent.com/adrianvlupu/C4-PlantUML/latest/C4_Dynamic.puml\n")
	//b.WriteString("LAYOUT_TOP_DOWN()\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	for _, element := range c.elements {
		b.WriteString(element.C4Writer())
	}

	for _, step := range c.sequence {
		b.WriteString(step.ToC4PlantUMLString())
	}

	b.WriteString("@enduml")

	return b.String()
}
