package diagrams

import (
	"fmt"
	"github.com/SteffiPeTaffy/c4documentation/elements"
	"testing"
)

func TestC4SequenceDiagram_ToC4PlantUMLString(t *testing.T) {
	someContainer := elements.
		NewContainer("my first container").
		Build()

	someOtherContainer := elements.
		NewContainer("my other service").
		Description("does also stuff").
		Technology("Go").
		RelatesTo(someContainer, "requests stuff", "REST/https")

	sequenceDiagram := NewSequenceDiagram("My Sequence Diagram").
	Next(*someContainer, *someOtherContainer, "Sends customer update events to", "async").
	Next(*someOtherContainer, *someContainer, "Sends ack back", "async")

	fmt.Println(sequenceDiagram.ToC4PlantUMLString())
}