package main

import (
	"c4documentation/diagrams"
	"c4documentation/elements"
	"fmt"
)

func main() {
	someContainer := elements.NewContainer("my first container", "does things", "Go")
	someContainerDatabase := elements.NewDatabase("my database", "stores stuff", "Postgres")
	someOtherContainer := func() *elements.Container {
		c := elements.NewContainer(
			"my other service",
			"does also stuff",
			"Go")
		c.RelatesTo(someContainer, "requests stuff", "REST/https")
		c.RelatesTo(someContainerDatabase, "persists stuff", "REST/https")
		return c
	}()
	someSystemBoundary := elements.NewSystemBoundary("boundary one", someContainer)
	someOtherSystemBoundary := elements.NewSystemBoundary("boundary two", someOtherContainer, someContainerDatabase)


	containerDiagram := diagrams.ContainerDiagram{
		Name: "SWF Container Diagram",
		Elements: []elements.PlantUMLElement{
			someSystemBoundary,
			someOtherSystemBoundary,
		},
	}

	fmt.Println(containerDiagram.ToPlantUMLString())
}
