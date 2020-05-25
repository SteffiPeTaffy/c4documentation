package elements

type C4Model struct {
	Elements []C4Element
}

func (m *C4Model) Contains(element C4Element) (found bool) {
	for _, elem := range m.Elements {
		if elem.Alias() == element.Alias() {
			found = true
		}
	}
	return found
}

func (m *C4Model) BuildBoundaryViewFrom(filter func(element C4Element) bool) *BoundaryView {
	baseBoundary := &BoundaryView{
		Parent:           *NewSystemBoundary("SWF").Build(),
		Children:         []C4Element{},
		NestedBoundaries: []BoundaryView{},
	}
	for _, element := range m.Elements {
		if filter(element) {
			if element.Parent == nil {
				baseBoundary.Children = append(baseBoundary.Children, element)
			} else {
				parentBoundary, hasParentBoundary := baseBoundary.FindParent(element)
				if !hasParentBoundary {
					baseBoundary.NestedBoundaries = append(baseBoundary.NestedBoundaries, BoundaryView{
						Parent:           *element.Parent,
						Children:         []C4Element{element},
						NestedBoundaries: []BoundaryView{},
					})
				} else {
					parentBoundary.Children = append(parentBoundary.Children, element)
				}
			}
		}
	}

	return baseBoundary
}
