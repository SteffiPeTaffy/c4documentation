package elements

import (
	"bytes"
	"fmt"
)

type BoundaryView struct {
	ElementInView    *SystemBoundary
	Children         []*C4Element
	NestedBoundaries []*BoundaryView
}

func NewBoundaryView(elements []*C4Element) *BoundaryView {
	root := &BoundaryView{
		ElementInView:    nil,
		Children:         []*C4Element{},
		NestedBoundaries: []*BoundaryView{},
	}

	// add elements that live on top level (not inside a boundary) as children
	withChilds := []*C4Element{}
	for _, element := range elements {
		if element.Parent == nil {
			root.Children = append(root.Children, element)
			continue
		}
		withChilds = append(withChilds, element)
	}
	// create list of all boundaries (and their children) that are referenced by one or more children
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
	// add all boundaries that are wrapping other boundaries and add their sub boundaries
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
	// all boundaries that are not contained within other boundaries live on the root
	for _, boundary := range boundaries {
		if boundary.ElementInView.Parent == nil {
			root.NestedBoundaries = append(root.NestedBoundaries, boundary)
		}
	}

	return root
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
		for _, layoutRelation := range nestedBoundary.ElementInView.LayoutRelations {
			buffer.WriteString(fmt.Sprintf("%s --[hidden]-- %s\n", layoutRelation.from.Alias(), layoutRelation.to.Alias()))
		}
	}

	if b.ElementInView != nil {
		buffer.WriteString("}\n")
	}

	return buffer.String()
}
