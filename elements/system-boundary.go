package elements

import (
	"bytes"
	"fmt"
)

type SystemBoundary struct {
	C4BoundaryElement
}

func NewSystemBoundary(name string) *SystemBoundary {
	systemBoundary := SystemBoundary{
		C4BoundaryElement: C4BoundaryElement{
			C4NodeElement: C4NodeElement{Name: name, OutgoingRelations: []C4Relation{}},
			elements:      []C4NodeElement{},
			containers:    []C4BoundaryElement{},
		},
	}
	systemBoundary.C4Writer = func() string {
		return systemBoundary.toC4PlantUMLString()
	}
	return &systemBoundary
}

func (sb *SystemBoundary) AddSystemBoundary(systemBoundary C4BoundaryElement) *SystemBoundary {
	sb.containers = append(sb.containers, systemBoundary)
	return sb
}

func (sb *SystemBoundary) Add(element C4NodeElement) *SystemBoundary {
	sb.elements = append(sb.elements, element)
	return sb
}

func (sb *SystemBoundary) toC4PlantUMLString() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", sb.Alias(), sb.Name))

	for _, element := range sb.elements {
		b.WriteString(element.C4Writer())
	}

	for _, container := range sb.containers {
		b.WriteString(container.C4Writer())
	}

	b.WriteString("}\n")

	return b.String()
}
