package elements

type SystemBoundary struct {
	C4BoundaryElement
}

func NewSystemBoundary(name string) *SystemBoundary {
	systemBoundary := SystemBoundary{
		C4BoundaryElement: C4BoundaryElement{
			C4BaseElement: C4BaseElement{
				Name:              name,
				OutgoingRelations: []C4Relation{},
			},
		},
	}
	return &systemBoundary
}
