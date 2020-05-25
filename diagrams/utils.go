package diagrams

import (
	"bytes"
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
)

func drawBoundary(boundary elements.BoundaryView) string {
	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", boundary.Parent.Alias(), boundary.Parent.Name))

	for _, element := range boundary.Children {
		b.WriteString(element.C4Writer())
	}

	for _, nestedBoundary := range boundary.NestedBoundaries {
		b.WriteString(drawBoundary(nestedBoundary))
	}

	b.WriteString("}\n")

	return b.String()
}

func drawBoundaryView(boundaryView *elements.BoundaryView) string {
	var b bytes.Buffer

	for _, element := range boundaryView.Children {
		b.WriteString(element.C4Writer())
	}

	for _, boundary := range boundaryView.NestedBoundaries {
		b.WriteString(drawBoundary(boundary))
	}
	return b.String()
}

func drawRelations(elements []elements.C4Element) string {
	var b bytes.Buffer

	for _, element := range elements {
		for _, relation := range element.OutgoingRelations {
			b.WriteString(relation.ToC4PlantUMLString())
		}
	}

	return b.String()
}


