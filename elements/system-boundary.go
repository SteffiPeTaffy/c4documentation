package elements

import (
	"bytes"
	"fmt"
)

type SystemBoundary struct {
	C4ContainerElement
}

func NewSystemBoundary(name string) *SystemBoundary {
	systemBoundary := SystemBoundary{
		C4ContainerElement: C4ContainerElement{
			C4NodeElement: C4NodeElement{Name: name, OutgoingRelations: []C4Relation{}},
			elements:      []C4NodeElement{},
			containers:    []C4ContainerElement{},
		},
	}
	systemBoundary.C4Writer = func() string {
		return systemBoundary.toC4PlantUMLString()
	}
	return &systemBoundary
}

func (sb *SystemBoundary) AddSystemBoundary(systemBoundary C4ContainerElement) *SystemBoundary {
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

	b.WriteString("}\n")

	return b.String()
}
