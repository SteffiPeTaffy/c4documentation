package diagrams

import (
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
	"testing"
)

func TestC4SequenceDiagram_ToC4PlantUMLString(t *testing.T) {
	someSystemBoundary := elements.
		NewSystemBoundary("boundary one").
		Build()

	someContainerDatabase := elements.
		NewDatabase("my database").
		Description("stores stuff").
		Technology("Postgres").
		BelongsTo(someSystemBoundary).
		Build()

	someContainer := elements.
		NewContainer("my first container").
		RelatesTo(someContainerDatabase, "persists stuff", "REST/https").
		BelongsTo(someSystemBoundary).
		Build()

	someOtherSystemBoundary := elements.
		NewSystemBoundary("boundary two").
		Build()

	someOtherContainer := elements.
		NewContainer("my other service").
		Description("does also stuff").
		Technology("Go").
		RelatesTo(someContainer, "requests stuff", "REST/https").
		BelongsTo(someOtherSystemBoundary).
		Build()

	myModel := &elements.C4Model{
		Elements: []*elements.C4Element{someContainer, someContainerDatabase, someOtherContainer},
	}

	sequenceDiagram := NewSequenceDiagram("My Sequence Diagram", myModel).
		Next(someContainer, someOtherContainer, "Sends customer update events to", "async").
		Next(someOtherContainer, someContainer, "Sends ack back", "async")

	fmt.Println(sequenceDiagram.ToC4PlantUMLString())
}
