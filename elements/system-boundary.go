package elements

type LayoutRelation struct {
	from *SystemBoundary
	to   *SystemBoundary
}

type SystemBoundary struct {
	*C4BaseElement
	LayoutRelations []*LayoutRelation
}

func NewSystemBoundary(name string) *SystemBoundary {
	return &SystemBoundary{
		C4BaseElement: &C4BaseElement{
			Name:              name,
			OutgoingRelations: []*C4Relation{},
		},
	}
}
func (systemBoundary *SystemBoundary) Build() *SystemBoundary {
	return systemBoundary
}

func (systemBoundary *SystemBoundary) BelongsTo(parent *SystemBoundary) *SystemBoundary {
	systemBoundary.Parent = parent
	return systemBoundary
}

func (systemBoundary *SystemBoundary) Above(elementBelow *SystemBoundary) *SystemBoundary {
	systemBoundary.LayoutRelations = append(systemBoundary.LayoutRelations, &LayoutRelation{
		from: systemBoundary,
		to:   elementBelow,
	})
	return systemBoundary
}

func (systemBoundary *SystemBoundary) Below(elementAbove *SystemBoundary) *SystemBoundary {
	systemBoundary.LayoutRelations = append(systemBoundary.LayoutRelations, &LayoutRelation{
		from: elementAbove,
		to:   systemBoundary,
	})
	return systemBoundary
}

func (systemBoundary *SystemBoundary) GetRoot(element *C4Element) (found *C4BaseElement) {
	if element.Parent == nil {
		found = element.C4BaseElement
	}

	return systemBoundary.GetRoot(element)
}
