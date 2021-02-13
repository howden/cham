package eval

import (
	"fmt"
	"github.com/howden/cham/ast"
)

type SimpleState struct {
	m map[ast.Identifier]int
}

func (s *SimpleState) GetVar(ident ast.Identifier) (int, error) {
	v, ok := s.m[ident]
	if ok {
		return v, nil
	} else {
		return 0, fmt.Errorf("no value for identifier %v", ident)
	}
}

func (s *SimpleState) PutVar(ident ast.Identifier, v int) {
	s.m[ident] = v
}

func (s *SimpleState) RemoveVar(ident ast.Identifier) {
	delete(s.m, ident)
}

func NewState() *SimpleState {
	return &SimpleState{make(map[ast.Identifier]int)}
}
