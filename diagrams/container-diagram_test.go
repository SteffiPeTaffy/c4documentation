package diagrams

import (
	"c4documentation/elements"
	"fmt"
	"testing"
)

func TestContainerDiagram_ToPlantUMLString(t *testing.T) {
	someContainer := elements.NewContainer("my first container", "does things", "Go").Build()
	someContainerDatabase := elements.NewDatabase("my database", "stores stuff", "Postgres").Build()
	someOtherContainer := elements.NewContainer(
			"my other service",
			"does also stuff",
			"Go").
		RelatesTo(&someContainer, "requests stuff", "REST/https").
		RelatesTo(&someContainerDatabase, "persists stuff", "REST/https").
		Build()

	someSystemBoundary := elements.NewSystemBoundary("boundary one", someContainer).Build()
	someOtherSystemBoundary := elements.NewSystemBoundary("boundary two", someOtherContainer, someContainerDatabase).Build()

	containerDiagram := ContainerDiagram{
		Name: "SWF Container Diagram",
		Elements: []elements.NamedElement{
			someSystemBoundary,
			someOtherSystemBoundary,
		},
	}

	fmt.Println(containerDiagram.ToPlantUMLString())
}