package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

type C4SequenceDiagram struct {
	name     string
	elements map[elements.C4Alias]elements.WritableElement
	sequence []*elements.Step
	model    *elements.C4Model
}

func NewSequenceDiagram(name string, model *elements.C4Model) *C4SequenceDiagram {
	return &C4SequenceDiagram{
		name:     name,
		model:    model,
		elements: map[elements.C4Alias]elements.WritableElement{},
	}
}

func (c *C4SequenceDiagram) Next(from elements.WritableElement, to elements.WritableElement, label string, dataObject string) *C4SequenceDiagram {
	c.elements[from.GetBase().Alias()] = from
	c.elements[to.GetBase().Alias()] = to
	c.sequence = append(c.sequence, &elements.Step{
		From:       from.GetBase(),
		To:         to.GetBase(),
		Label:      label,
		DataObject: dataObject,
	})
	return c
}

func (c *C4SequenceDiagram) ToC4PlantUMLString() string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("@startuml %s\n", c.name))
	b.WriteString("!include https://raw.githubusercontent.com/SteffiPeTaffy/c4documentation/master/templates/sequence.puml\n")
	b.WriteString("LAYOUT_TOP_DOWN()\n")
	b.WriteString("LAYOUT_WITH_LEGEND()\n")

	onlyRelevant := func(element elements.WritableElement) bool {
		_, ok := c.elements[element.GetBase().Alias()]
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
