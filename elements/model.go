package elements

type BoundaryView struct {
	Parent           C4BoundaryElement
	Children         []C4Element
	NestedBoundaries []BoundaryView
}

func (b *BoundaryView) VisitBoundaries(callback func(parent BoundaryView) (done bool)) {
	if done := callback(*b); done {
		return
	}

	for _, nestedBoundary := range b.NestedBoundaries {
		nestedBoundary.VisitBoundaries(callback)
	}
}

func (b *BoundaryView) FindParent(child C4Element) (foundParent *BoundaryView, found bool) {
	if child.Parent == nil {
		return nil, false
	}

	b.VisitBoundaries(func(elem BoundaryView) (done bool) {
		if elem.Parent.Alias() == child.Parent.Alias() {
			found = true
			foundParent = &elem
			return true
		}
		return false
	})
	return
}

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

func (m *C4Model) BuildBoundaryViewFrom(elements []C4Element) *BoundaryView {
	baseBoundary := &BoundaryView{
		Parent:           *NewSystemBoundary("SWF").Build(),
		Children:         []C4Element{},
		NestedBoundaries: []BoundaryView{},
	}
	for _, element := range elements {
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

	return baseBoundary
}
