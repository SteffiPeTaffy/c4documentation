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
	return &Container {
		NamedElement: NamedElement{Name: name},
		Description:  description,
		Technology:   technology,
	}
}

func (c *Container) ToPlantUMLString() string {
	return fmt.Sprintf("Container(%v, '%s', '%s', '%s')\n", c.Alias(), c.Name, c.Technology, c.Description)
}