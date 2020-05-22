package diagrams

import (
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
	"testing"
)

func TestC4SequenceDiagram_ToC4PlantUMLString(t *testing.T) {
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

	boundaryWrappingSomeOtherBonudary := elements.
		NewSystemBoundary("wrapping another boundary").
		AddSystemBoundary(*someOtherSystemBoundary).
		Build()

	myModel := elements.C4Model{
		Boundaries: []elements.C4BoundaryElement{*someSystemBoundary, *boundaryWrappingSomeOtherBonudary},
	}

	sequenceDiagram := NewSequenceDiagram("My Sequence Diagram", myModel).
	Next(*someContainer, *someOtherContainer, "Sends customer update events to", "async").
	Next(*someOtherContainer, *someContainer, "Sends ack back", "async")

	fmt.Println(sequenceDiagram.ToC4PlantUMLString())
}