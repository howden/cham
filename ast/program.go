package ast

import "fmt"

type Program struct {
	Input     []int
	Reactions []*Reaction
}

func (program Program) String() string {
	return fmt.Sprintf("input{\n  %v\n}\n%v", program.Input, program.Reactions)
}
