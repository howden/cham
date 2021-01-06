package eval

import "github.com/howden/cham/ast"

type SimpleState struct {
	m map[ast.Identifier]int
}

func (s *SimpleState) GetVar(ident ast.Identifier) int {
	return s.m[ident]
}

func (s *SimpleState) PutVar(ident ast.Identifier, v int) {
	s.m[ident] = v
}

func NewState() *SimpleState {
	return &SimpleState{make(map[ast.Identifier]int)}
}
