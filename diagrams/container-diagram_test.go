package diagrams

import (
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
	"testing"
)

func TestContainerDiagram_ToPlantUMLString(t *testing.T) {
	someContainerDatabase := elements.
		NewDatabase("my database").
		Description("stores stuff").
		Technology("Postgres").
		Build()
	someContainer := elements.
		NewContainer("my first container").
		RelatesTo(someContainerDatabase, "persists stuff", "REST/https")

	someSystemBoundary := elements.
		NewSystemBoundary("boundary one").
		Add(*someContainer).
		Add(*someContainerDatabase).
		Build()

	someOtherContainer := elements.
		NewContainer("my other service").
		Description("does also stuff").
		Technology("Go").
		RelatesTo(someContainer, "requests stuff", "REST/https")

	someOtherSystemBoundary := elements.
		NewSystemBoundary("boundary two").
		Add(*someOtherContainer).
		Build()

	containerDiagram := NewContainerDiagram("SWF Container Diagram").
		AddSystemBoundary(*someSystemBoundary).
		AddSystemBoundary(*someOtherSystemBoundary)


	fmt.Println(containerDiagram.ToC4PlantUMLString())
}
