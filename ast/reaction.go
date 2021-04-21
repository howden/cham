package ast

import (
	"fmt"
)

// AST encapsulating a defined reaction
type ReactionPointer struct {
	Identifier Identifier
	Reactions  []*Reaction
}

// AST encapsulating a single reaction
type Reaction struct {
	Input     *ReactionInput
	Action    *ReactionAction
	Condition *ReactionCondition
}

// Represents the reaction input
// Just an array of identifier tuples, can be empty.
type ReactionInput struct {
	Idents []IdentifierTuple
}

// Represents the reaction action
// This is formed of products: an array of int term tuples, which could be empty.
type ReactionAction struct {
	Products []IntegerTermTuple
}

// Represents the reaction condition
// Just a single boolean expression (of course, this can
// be expanded infinitely with and/or rules)
type ReactionCondition struct {
	Expression BooleanTerm
}

func (reaction Reaction) String() string {
	f := `reaction{
  input{
    %v
  },
  action{
    %v
  },
  condition{
    %v
  },
}`
	return fmt.Sprintf(f, reaction.Input.Idents, reaction.Action.Products, reaction.Condition.Expression)
}
