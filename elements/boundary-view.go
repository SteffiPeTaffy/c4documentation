package elements

import (
	"bytes"
	"fmt"
)

type BoundaryView struct {
	ElementInView    *C4BoundaryElement
	Children         []*C4Element
	NestedBoundaries []*BoundaryView
}

func NewBoundaryView(elements []*C4Element) *BoundaryView {
	root := &BoundaryView{
		ElementInView:    nil,
		Children:         []*C4Element{},
		NestedBoundaries: []*BoundaryView{},
	}

	withChilds := []*C4Element{}
	for _, element := range elements {
		if element.Parent == nil {
			root.Children = append(root.Children, element)
			continue
		}
		withChilds = append(withChilds, element)
	}
	boundaries := make(map[C4Alias]*BoundaryView)
	for _, element := range withChilds {
		if boundaries[element.Parent.Alias()] != nil {
			boundaries[element.Parent.Alias()].Children = append(boundaries[element.Parent.Alias()].Children, element)
		} else {
			boundaries[element.Parent.Alias()] = &BoundaryView{
				ElementInView:    element.Parent,
				Children:         []*C4Element{element},
				NestedBoundaries: []*BoundaryView{},
			}
		}
	}

	for _, boundary := range boundaries {
		p := boundary.ElementInView.Parent
		for p != nil {
			if boundaries[p.Alias()] != nil {
				boundaries[p.Alias()].NestedBoundaries = append(boundaries[p.Alias()].NestedBoundaries, boundary)
			} else {
				boundaries[p.Alias()] = &BoundaryView{
					ElementInView:    p,
					Children:         []*C4Element{},
					NestedBoundaries: []*BoundaryView{boundary},
				}
			}
			p = p.Parent
		}
	}
	for _, boundary := range boundaries {
		if boundary.ElementInView.Parent == nil {
			root.NestedBoundaries = append(root.NestedBoundaries, boundary)
		}
	}

	return root
}

func (b *BoundaryView) VisitBoundaries(callback func(parent *BoundaryView) (done bool)) {
	if done := callback(b); done {
		return
	}

	for _, nestedBoundary := range b.NestedBoundaries {
		nestedBoundary.VisitBoundaries(callback)
	}
}

func (b *BoundaryView) FindParent(child *C4Element) (foundParent *BoundaryView, found bool) {
	if child.Parent == nil {
		return nil, false
	}

	b.VisitBoundaries(func(elem *BoundaryView) (done bool) {
		if elem.ElementInView.Alias() == child.Parent.Alias() {
			found = true
			foundParent = elem
			return true
		}
		return false
	})
	return
}

func (b *BoundaryView) ToC4PlantUMLString() string {
	var buffer bytes.Buffer

	if b.ElementInView != nil {
		buffer.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", b.ElementInView.Alias(), b.ElementInView.Name))
	}
	for _, element := range b.Children {
		buffer.WriteString(element.C4Writer())
	}

	for _, nestedBoundary := range b.NestedBoundaries {
		buffer.WriteString(nestedBoundary.ToC4PlantUMLString())
	}

	if b.ElementInView != nil {
		buffer.WriteString("}\n")
	}
	return buffer.String()
}

/*
func boundaryToC4PlantUMLString(boundary *BoundaryView) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("System_Boundary(%s, %s) {\n", boundary.Parent.Alias(), boundary.Parent.Name))

	for _, element := range boundary.Children {
		buffer.WriteString(element.C4Writer())
	}

	for _, nestedBoundary := range boundary.NestedBoundaries {
		buffer.WriteString(nestedBoundary.ToC4PlantUMLString())
	}

	buffer.WriteString("}\n")

	return buffer.String()
}*/
