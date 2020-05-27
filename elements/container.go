package elements

import (
	"fmt"
)

type Container struct {
	*C4Element
	description string
	owner       string
}

func NewContainer(name string) *Container {
	container := Container{
		C4Element:   &C4Element{
			C4BaseElement:   &C4BaseElement{
				Name:              name,
				OutgoingRelations: []*C4Relation{},
			},
		},
	}
	container.C4Writer = container.toC4PlantUMLString
	return &container
}

func (c *Container) Description(description string) *Container {
	c.description = description
	return c
}

func (c *Container) Owner(owner string) *Container {
	c.owner = owner
	return c
}

func (c *Container) toC4PlantUMLString() string {
	return fmt.Sprintf("Container(%v, '%s', '%s', '%s')\n", c.Alias(), c.Name, c.owner, c.description)
}
