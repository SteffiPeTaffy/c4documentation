package elements

import (
	"fmt"
)

type Relation struct {
	From       NamedElement
	To         AliasElement
	Label      string
	Technology string
}

func (r *Relation) ToPlantUMLString() string {
	return fmt.Sprintf("Rel(%v, '%v', '%s', '%s')\n", r.From.Alias(), r.To.Alias(), r.Label, r.Technology)
}