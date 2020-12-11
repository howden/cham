package ast

// AST encapsulating a single reaction
type Reaction struct {
	Input     ReactionInput
	Action    ReactionAction
	Condition ReactionCondition
}

// Represents the reaction input
// Just an array of identifiers, can be empty.
type ReactionInput struct {
	Idents []Identifier
}

// Represents the reaction action
// This is formed of products: an array of int terms, which could be empty.
type ReactionAction struct {
	Products []IntegerTerm
}

// Represents the reaction condition
// Just a single boolean expression (of course, this can
// be expanded infinitely with and/or rules)
type ReactionCondition struct {
	Expression BooleanTerm
}
