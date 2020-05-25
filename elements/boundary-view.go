package elements

import (
	"bytes"
	"fmt"
)

type BoundaryView struct {
	Parent           C4BoundaryElement
	Children         []C4Element
	NestedBoundaries []BoundaryView
}

func (b *BoundaryView) VisitBoundaries(callback func(parent BoundaryView) (done bool)) {
	if done := callback(*b); done {
		return
	}

	for _, nestedBoundary := range b.NestedBoundaries {
		nestedBoundary.VisitBoundaries(callback)
	}
}

func (b *BoundaryView) FindParent(child C4Element) (foundParent *BoundaryView, found bool) {
	if child.Parent == nil {
		return nil, false
	}

	b.VisitBoundaries(func(elem BoundaryView) (done bool) {
		if elem.Parent.Alias() == child.Parent.Alias() {
			found = true
			foundParent = &elem
			return true
		}
		return false
	})
	return
}

func (b *BoundaryView) ToC4PlantUMLString() string {
	var buffer bytes.Buffer

	for _, element := range b.Children {
		buffer.WriteString(element.C4Writer())
	}

	for _, nestedBoundary := range b.NestedBoundaries {
		buffer.WriteString(boundaryToC4PlantUMLString(nestedBoundary))
	}

	return buffer.String()
}

func boundaryToC4PlantUMLString(boundary BoundaryView) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", boundary.Parent.Alias(), boundary.Parent.Name))

	for _, element := range boundary.Children {
		buffer.WriteString(element.C4Writer())
	}

	for _, nestedBoundary := range boundary.NestedBoundaries {
		buffer.WriteString(nestedBoundary.ToC4PlantUMLString())
	}

	buffer.WriteString("}\n")

	return buffer.String()
}
