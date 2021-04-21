package eval

import (
	"fmt"
	"github.com/howden/cham/ast"
)

type ReactionStore struct {
	m map[ast.Identifier][]*ast.Reaction
}

func (s *ReactionStore) Get(ident ast.Identifier) ([]*ast.Reaction, error) {
	v, ok := s.m[ident]
	if ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("no value for identifier %v", ident)
	}
}

func (s *ReactionStore) Put(def *ast.ReactionPointer) {
	s.m[def.Identifier] = def.Reactions
}

func (s *ReactionStore) Delete(ident ast.Identifier) {
	delete(s.m, ident)
}

func (s *ReactionStore) Slice() []*ast.ReactionPointer {
	res := make([]*ast.ReactionPointer, 0, len(s.m))
	for ident, reactions := range s.m {
		res = append(res, &ast.ReactionPointer{
			Identifier: ident,
			Reactions:  reactions,
		})
	}
	return res
}

func NewReactionStore() *ReactionStore {
	return &ReactionStore{make(map[ast.Identifier][]*ast.Reaction)}
}
