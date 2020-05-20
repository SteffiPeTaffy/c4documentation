package elements

import (
	"bytes"
	"fmt"
)

type SystemBoundary struct {
	NamedElement
	Elements []NamedElement
}

func NewSystemBoundary(name string, elements ...NamedElement) *SystemBoundary {
	systemBoundary := SystemBoundary{
		NamedElement: NamedElement{
			Name: name,
		},
		Elements: elements,
	}
	systemBoundary.C4Writer = systemBoundaryWriter(systemBoundary)
	return &systemBoundary
}

func systemBoundaryWriter(sb SystemBoundary) func(element *NamedElement) string {
	return func(element *NamedElement) string {
		var b bytes.Buffer
		b.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", sb.Alias(), sb.Name))

		for _, element := range sb.Elements {
			b.WriteString(element.C4Writer(&element))
		}

		b.WriteString("}\n")

		return b.String()
	}
}



func (sb *SystemBoundary) Add(element NamedElement) *SystemBoundary {
	sb.Elements = append(sb.Elements, element)
	return sb
}