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
	database := Database{
		NamedElement: NamedElement{Name: name},
		Description:  description,
		Technology:   technology,
	}
	database.C4Writer = DatabaseWriter(database)
	return &database
}

func DatabaseWriter(d Database) func(element *NamedElement) string {
	return func(element *NamedElement) string {
		return fmt.Sprintf("ContainerDb(%s, %s, %s, %s)\n", d.Alias(), d.Name, d.Technology, d.Description)
	}
}
