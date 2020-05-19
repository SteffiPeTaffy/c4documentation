package elements

import (
	"fmt"
)

type Database struct {
	NamedElement
	Description string
	Technology  string
}

func NewDatabase(name string, description string, technology string) *Database {
	return &Database {
		NamedElement: NamedElement{Name: name},
		Description:  description,
		Technology:   technology,
	}
}

func (d *Database) ToPlantUMLString() string {
	return fmt.Sprintf("ContainerDb(%s, %s, %s, %s)\n", d.Alias(), d.Name, d.Technology, d.Description)
}
