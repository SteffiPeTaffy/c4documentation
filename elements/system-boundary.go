package elements

import (
	"bytes"
	"fmt"
)

type SystemBoundary struct {
	NamedElement
	Elements []PlantUMLElement
}

func NewSystemBoundary(name string, elements ...PlantUMLElement) *SystemBoundary {
	return &SystemBoundary {
		NamedElement: NamedElement{Name: name},
		Elements: elements,
	}
}

func (sb *SystemBoundary) Add(element PlantUMLElement) *SystemBoundary {
	sb.Elements = append(sb.Elements, element)
	return sb
}

func (sb *SystemBoundary) ToPlantUMLString() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", sb.Alias(), sb.Name))

	for _, element := range sb.Elements {
		b.WriteString(element.ToPlantUMLString())
	}

	b.WriteString("}\n")

	return b.String()
}