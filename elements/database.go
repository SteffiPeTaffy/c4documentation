package elements

import (
	"fmt"
)

type Database struct {
	*C4Element
	description string
	owner       string
}

func NewDatabase(name string) *Database {
	database := Database{
		C4Element:   &C4Element{
			C4BaseElement:   &C4BaseElement{
				Name:              name,
				OutgoingRelations: []*C4Relation{},
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

func (d *Database) Owner(owner string) *Database {
	d.owner = owner
	return d
}

func (d *Database) BelongsTo(parent *SystemBoundary) *Database {
	d.C4Element.BelongsTo(parent)
	return d
}

func (d *Database) RelatesTo(to ElementWithBase, label string, technology string) *Database {
	d.C4Element.RelatesTo(to,label,technology)
	return d
}

func (d *Database) toC4PlantUMLString() string {
	return fmt.Sprintf("ContainerDb(%v, '%s', '%s', '%s')\n", d.Alias(), d.Name, d.owner, d.description)
}
