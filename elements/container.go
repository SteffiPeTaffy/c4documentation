package elements

import (
	"fmt"
)

type Container struct {
	C4NodeElement
	description       string
	technology        string
}

func NewContainer(name string) *Container {
	container := Container{
		C4NodeElement: C4NodeElement{Name: name, OutgoingRelations: []C4Relation{}},
	}
	container.C4Writer = func() string {
		return container.toC4PlantUMLString()
	}
	return &container
}

func (c *Container) Description(description string) *Container {
	c.description = description
	return c
}

func (c *Container) Technology(technology string) *Container {
	c.technology = technology
	return c
}

func (c *Container) toC4PlantUMLString() string {
	return fmt.Sprintf("Container(%v, '%s', '%s', '%s')\n", c.Alias(), c.Name, c.technology, c.description)
}
