package ast

import "fmt"

// A tuple is like an array, but with a shape that is known at 'compile' time.
// It is basically a composite value holder for zero or more ints, int terms or identifiers.
// Tuples are limited to one dimension - a tuple cannot contain another tuple (can't be nested)
type Tuple interface {
	Dimensions() int
}

// Checks if the dimensions of two tuples match
func ShapeMatches(a Tuple, b Tuple) bool {
	return a.Dimensions() == b.Dimensions()
}

// A tuple of IntegerTerms
type IntegerTermTuple struct {
	Values []IntegerTerm
}

func (tuple IntegerTermTuple) Dimensions() int {
	return len(tuple.Values)
}

func (tuple *IntegerTermTuple) String() string {
	return fmt.Sprintf("intTermTuple(%s)", tuple.Values)
}

// A tuple of Identifiers
type IdentifierTuple struct {
	Values []Identifier
}

func (tuple IdentifierTuple) Dimensions() int {
	return len(tuple.Values)
}

func (tuple *IdentifierTuple) String() string {
	return fmt.Sprintf("identTuple(%s)", tuple.Values)
}

// A tuple of ints
// As this struct is used as a map key, we cannot use a slice. Therefore the shape is limited to 1x16.
// 16 was chosen arbitrarily, it could be increased if needed.
type IntTuple struct {
	Shape  int
	Values [16]int
}

func (tuple IntTuple) Slice() []int {
	return tuple.Values[0:tuple.Shape]
}

func (tuple IntTuple) Dimensions() int {
	return tuple.Shape
}

func (tuple IntTuple) String() string {
	if tuple.Shape == 1 {
		return fmt.Sprint(tuple.Values[0])
	}
	return fmt.Sprint(tuple.Slice())
}

func CreateIntTuple(values []int) IntTuple {
	shape := len(values)
	if shape > 16 {
		shape = 16
	}
	if shape < 1 {
		panic("shape cannot be 0")
	}

	var arr [16]int
	copy(arr[:], values[0:shape])

	return IntTuple{
		Shape:  shape,
		Values: arr,
	}
}
