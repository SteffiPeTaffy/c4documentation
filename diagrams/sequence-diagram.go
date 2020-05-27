package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

type C4SequenceDiagram struct {
	name     string
	elements map[elements.C4Alias]*elements.C4Element
	sequence []*elements.Step
	model    *elements.C4Model
}

func NewSequenceDiagram(name string, model *elements.C4Model) *C4SequenceDiagram {
	return &C4SequenceDiagram{
		name:     name,
		model:    model,
		elements: map[elements.C4Alias]*elements.C4Element{},
	}
}

func (c *C4SequenceDiagram) Next(from *elements.C4Element, to *elements.C4Element, label string, dataObject string) *C4SequenceDiagram {
	c.elements[from.Alias()] = from
	c.elements[to.Alias()] = to
	c.sequence = append(c.sequence, &elements.Step{
		From:       from,
		To:         to,
		Label:      label,
		DataObject: dataObject,
	})
	return c
}

func (c *C4SequenceDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!include https://raw.githubusercontent.com/adrianvlupu/C4-PlantUML/latest/C4_Dynamic.puml\n")
	b.WriteString("LAYOUT_TOP_DOWN()\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	onlyRelevant := func(element *elements.C4Element) bool {
		_, ok := c.elements[element.Alias()]
		return ok
	}
	b.WriteString(c.model.CreateBoundaryView(onlyRelevant).ToC4PlantUMLString())

	for _, step := range c.sequence {
		if c.model.Contains(step.From) || c.model.Contains(step.To) {
			b.WriteString(step.ToC4PlantUMLString())
		}
	}

	b.WriteString("@enduml")

	return b.String()
}
