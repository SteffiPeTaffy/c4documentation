package elements

import (
	"fmt"
)

type MessageQueue struct {
	*C4Element
	description string
	owner       string
}

func NewMessageQueue(name string) *MessageQueue {
	database := MessageQueue{
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

func (d *MessageQueue) Description(description string) *MessageQueue {
	d.description = description
	return d
}

func (d *MessageQueue) Owner(owner string) *MessageQueue {
	d.owner = owner
	return d
}

func (d *MessageQueue) toC4PlantUMLString() string {
	return fmt.Sprintf("ContainerQ(%v, '%s', '%s', '%s')\n", d.Alias(), d.Name, d.owner, d.description)
}
