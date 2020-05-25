package elements

type C4Model struct {
	Elements []*C4Element
}

func (m *C4Model) Contains(element *C4Element) (found bool) {
	for _, elem := range m.Elements {
		if elem.Alias() == element.Alias() {
			found = true
		}
	}
	return found
}

func (m *C4Model) CreateBoundaryView(filter func(element *C4Element) bool) *BoundaryView {
	var filteredElements []*C4Element

	for _, elem := range m.Elements {
		if filter(elem) {
			filteredElements = append(filteredElements, elem)
		}
	}

	return NewBoundaryView(filteredElements)
}
