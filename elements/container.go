package elements

import (
	"fmt"
)

type Container struct {
	NamedElement
	Description string
	Technology  string
}

func NewContainer(name string, description string, technology string) *Container {
	container := Container{
		NamedElement: NamedElement{Name: name},
		Description:  description,
		Technology:   technology,
	}
	container.C4Writer = containerWriter(container)
	return &container
}

func containerWriter(c Container) func(n *NamedElement) string {
	return func(n *NamedElement) string {
		return fmt.Sprintf("Container(%v, '%s', '%s', '%s')\n", c.Alias(), c.Name, c.Technology, c.Description)
	}
}