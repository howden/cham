package analysis

import "github.com/howden/cham/ast"

// Represents the statically determined reaction type.
type ReactionType int

const (
	// Represents a reaction that reduces/shrinks the size of the solution
	Reducing ReactionType = iota
	// Represents a reaction that keeps the size of the solution constant
	Constant
	// Represents a reaction that expands the size of the solution
	Expanding
)

// Determines the type of the reaction based on the number of inputs and outputs it has
func DetermineReactionType(reaction *ast.Reaction) ReactionType {
	inputs := len(reaction.Input.Idents)
	outputs := len(reaction.Action.Products)

	if inputs > outputs {
		return Reducing
	} else if outputs > inputs {
		return Expanding
	} else {
		return Constant
	}
}
