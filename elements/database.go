package elements

import (
	"fmt"
)

type Database struct {
	C4Element
	description       string
	technology        string
}

func NewDatabase(name string) *Database {
	database := Database{
		C4Element:   C4Element{
			C4BaseElement:   C4BaseElement{
				Name:              name,
				OutgoingRelations: []C4Relation{},
			},
		},
	}
	database.C4Writer = func() string {
		return database.toC4PlantUMLString()
	}
	return &database
}

func (d *Database) Description(description string) *Database {
	d.description = description
	return d
}

func (d *Database) Technology(technology string) *Database {
	d.technology = technology
	return d
}

func (d *Database) toC4PlantUMLString() string {
	return fmt.Sprintf("ContainerDb(%v, '%s', '%s', '%s')\n", d.Alias(), d.Name, d.technology, d.description)
}
