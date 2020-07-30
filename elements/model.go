package elements

type C4Model struct {
	Elements []*C4Element
}

func (m *C4Model) Contains(element C4PlantUMLAlias) bool {
	for _, elem := range m.Elements {
		if elem.Alias() == element.Alias() {
			return true
		}
	}
	return false
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

func (m *C4Model) Filter(f func(e *C4Element) bool) *C4Model {
	filteredElements := make([]*C4Element, len(m.Elements))
	for i := range m.Elements {
		if f(m.Elements[i]) {
			filteredElements[i] = m.Elements[i]
		}
	}
	return &C4Model{
		Elements: filteredElements,
	}
}
