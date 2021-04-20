package ast

import (
	"fmt"
)

// A program is made up of input in the form of an array of int tuples,
// and a sequence of reaction rules to apply in order.
type Program struct {
	Input     []IntTuple
	Reactions []*Reaction
}

func (program Program) String() string {
	return fmt.Sprintf("input{\n  %v\n}\n%v", program.Input, program.Reactions)
}
