package elements

import "fmt"

type C4Relation struct {
	From       *C4BaseElement
	To         C4PlantUMLAlias
	Label      string
	Technology string
}

func (r *C4Relation) ToC4PlantUMLString() string {
	return fmt.Sprintf("Rel(%v, '%v', '%s', '%s')\n", r.From.Alias(), r.To.Alias(), r.Label, r.Technology)
}
