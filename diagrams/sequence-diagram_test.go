package diagrams

import (
	"fmt"
	"testing"
)

func TestC4SequenceDiagram_ToC4PlantUMLString(t *testing.T) {
	sequenceDiagram := NewSequenceDiagram("My Sequence Diagram", myModel).
		Next(someContainer, someOtherContainer, "Sends customer update events to", "async").
		Next(someOtherContainer, someContainer, "Sends ack back", "async")

	fmt.Println(sequenceDiagram.ToC4PlantUMLString())
}
