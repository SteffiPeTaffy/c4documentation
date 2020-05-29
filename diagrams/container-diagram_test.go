package diagrams

import (
	"fmt"
	"testing"
)

func TestContainerDiagram_ToPlantUMLString(t *testing.T) {
	containerDiagram := NewContainerDiagram("My Container Diagram", myModel)

	fmt.Println(containerDiagram.ToC4PlantUMLString())
}
