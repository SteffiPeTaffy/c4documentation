package diagrams

import (
	"fmt"
	"testing"
)

func TestC4ContextDiagram_ToC4PlantUMLString(t *testing.T) {
	containerDiagram := NewContextDiagram("My Context Diagram", myModel)

	fmt.Println(containerDiagram.ToC4PlantUMLString())
}
