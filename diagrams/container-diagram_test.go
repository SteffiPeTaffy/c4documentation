package diagrams

import (
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
	"testing"
)

func TestContainerDiagram_ToPlantUMLString(t *testing.T) {
	someSystemBoundary := elements.
		NewSystemBoundary("boundary one").
		Build()

	someContainerDatabase := elements.
		NewDatabase("my database").
		Description("stores stuff").
		Technology("Postgres").
		BelongsTo(*someSystemBoundary).
		Build()

	someContainer := elements.
		NewContainer("my first container").
		RelatesTo(someContainerDatabase, "persists stuff", "REST/https").
		BelongsTo(*someSystemBoundary).
		Build()

	someOtherSystemBoundary := elements.
		NewSystemBoundary("boundary two").
		Build()

	someOtherContainer := elements.
		NewContainer("my other service").
		Description("does also stuff").
		Technology("Go").
		RelatesTo(someContainer, "requests stuff", "REST/https").
		BelongsTo(*someOtherSystemBoundary).
		Build()

	myModel := elements.C4Model{
		Elements: []elements.C4Element{*someContainer, *someContainerDatabase, *someOtherContainer},
	}

	containerDiagram := NewContainerDiagram("My Container Diagram", myModel)

	fmt.Println(containerDiagram.ToC4PlantUMLString())
}
