package elements

import "fmt"

type Step struct {
	C4Printable
	From       C4Element
	To         C4Element
	Label      string
	Technology string
}

func (s *Step) ToC4PlantUMLString() string {
	return fmt.Sprintf("Rel(%v, '%v', '%s', '%s')\n", s.From.Alias(), s.To.Alias(), s.Label, s.Technology)
}