package elements

type C4Model struct {
	Elements   []C4NodeElement
	Boundaries []C4BoundaryElement
}

func (m C4Model) Contains(element C4NodeElement) (found bool) {
	found = false
	for _, boundary := range m.Boundaries {
		boundary.VisitElements(func(elem C4NodeElement) (done bool) {
			if elem.Alias() == element.Alias() {
				found = true
				return true
			}
			return false
		})
	}

	for _, elem := range m.Elements {
		if elem.Alias() == element.Alias() {
			found = true
		}
	}
	return found
}