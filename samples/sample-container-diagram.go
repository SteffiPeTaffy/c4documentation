package main

import (
	"c4documentation/diagrams"
	"c4documentation/elements"
	"fmt"
)

func main() {
	someContainerDatabase := elements.
		NewDatabase("my database").
		Description("stores stuff").
		Technology("Postgres").
		Build()
	someContainer := elements.
		NewContainer("my first container").
		RelatesTo(*someContainerDatabase, "persists stuff", "REST/https")

	someSystemBoundary := elements.
		NewSystemBoundary("boundary one").
		Add(*someContainer).
		Add(*someContainerDatabase).
		Build()

	someOtherContainer := elements.
		NewContainer("my other service").
		Description("does also stuff").
		Technology("Go").
		RelatesTo(*someContainer, "requests stuff", "REST/https")

	someOtherSystemBoundary := elements.
		NewSystemBoundary("boundary two").
		Add(*someOtherContainer).
		Build()

	containerDiagram := diagrams.NewContainerDiagram("SWF Container Diagram").
		AddSystemBoundary(*someSystemBoundary).
		AddSystemBoundary(*someOtherSystemBoundary)


	fmt.Println(containerDiagram.ToC4PlantUMLString())
}
