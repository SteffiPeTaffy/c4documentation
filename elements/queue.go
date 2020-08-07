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

func (q *MessageQueue) Description(description string) *MessageQueue {
	q.description = description
	return q
}

func (q *MessageQueue) Owner(owner string) *MessageQueue {
	q.owner = owner
	return q
}

func (q *MessageQueue) BelongsTo(parent *SystemBoundary) *MessageQueue {
	q.C4Element.BelongsTo(parent)
	return q
}

func (q *MessageQueue) RelatesTo(to ElementWithBase, label string, technology string) *MessageQueue {
	q.C4Element.RelatesTo(to,label,technology)
	return q
}

func (q *MessageQueue) toC4PlantUMLString() string {
	return fmt.Sprintf("ContainerQ(%v, '%s', '%s', '%s')\n", q.Alias(), q.Name, q.owner, q.description)
}
