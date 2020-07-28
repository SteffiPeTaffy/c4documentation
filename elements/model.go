package elements

type WritableElement interface {
	ElementWithBase
	WritePUML() string
}

type ElementWithBase interface {
	GetBase() *C4BaseElement
}

type C4Model struct {
	Elements []WritableElement
}

func (m *C4Model) Contains(element C4PlantUMLAlias) bool {
	for _, elem := range m.Elements {
		if elem.GetBase().Alias() == element.Alias() {
			return true
		}
	}
	return false
}

func (m *C4Model) CreateBoundaryView(filter func(element WritableElement) bool) *BoundaryView {
	return NewBoundaryView(m.Filter(filter).Elements)
}

func (m *C4Model) Filter(f func(e WritableElement) bool) *C4Model {
	filteredElements := make([]WritableElement, 0)
	for i := range m.Elements {
		if f(m.Elements[i]) {
			filteredElements = append(filteredElements, m.Elements[i])
		}
	}
	return &C4Model{
		Elements: filteredElements,
	}
}
