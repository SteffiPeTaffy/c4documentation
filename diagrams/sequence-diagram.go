package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

type C4SequenceDiagram struct {
	elements.C4Printable
	name     string
	elements []elements.C4Element
	sequence []elements.Step
	model    elements.C4Model
}

func NewSequenceDiagram(name string, model elements.C4Model) *C4SequenceDiagram {
	return &C4SequenceDiagram{
		name:     name,
		model:    model,
	}
}

func (c *C4SequenceDiagram) Next(from elements.C4Element, to elements.C4Element, label string, technology string) *C4SequenceDiagram {
	c.elements = append(c.elements, from, to)
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
	b.WriteString("LAYOUT_TOP_DOWN()\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	relevantElements := make([]elements.C4Element, 0, len(c.elements))
	for  _, value := range c.elements {
		relevantElements = append(relevantElements, value)
	}

	b.WriteString(drawBoundaryView(c.model.BuildBoundaryViewFrom(relevantElements)))

	for _, step := range c.sequence {
		if c.model.Contains(step.From) || c.model.Contains(step.To) {
			b.WriteString(step.ToC4PlantUMLString())
		}
	}

	b.WriteString("@enduml")

	return b.String()
}
