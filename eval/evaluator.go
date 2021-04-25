package eval

import (
	"github.com/howden/cham/ast"
	"github.com/pkg/errors"
	"gonum.org/v1/gonum/stat/combin"
)

// Function to evaluate a program.
func Evaluate(prog *ast.Program) (*Multiset, error) {
	// Create a new multiset containing the program input
	multiset := NewMultiset()
	multiset.AddAll(prog.Input)

	for _, reaction := range prog.Reactions {
		err := evaluateReaction(reaction, multiset)
		if err != nil {
			return nil, errors.Wrap(err, "error evaluating reaction")
		}
	}

	return multiset, nil
}

// Function to evaluate a reaction.
func evaluateReaction(prog *ast.Reaction, multiset *Multiset) error {

	// Obtain a list of the identifiers used by the (single) reaction rule
	// The length of this array becomes 'k' in the k-permutations calculation
	k := len(prog.Input.Idents)

	// Continuously attempt reactions until either:
	// - 'n < k' where n is the cardinality of the multiset and k is the number of identifiers
	//     i.e. there's more variables in the reaction than there are values to fill them
	// - a previous iteration of the loop was unable to complete a single reaction
	solved := false
	for !solved && multiset.Cardinality() >= k {
		didReactionOccur, err := attemptReaction(prog, k, multiset)
		if err != nil {
			return err
		}

		// If a reaction didn't take place during this iteration, then set solved=true
		if !didReactionOccur {
			solved = true
		}
	}

	return nil
}

// Attempts to perform a single reaction within the multiset (solution).
//
// If/when a reaction takes place, the function will return immediately (with the value true).
// If after trying using all possible permutations a reaction has not taken place, the function will return false.
func attemptReaction(prog *ast.Reaction, k int, multiset *Multiset) (bool, error) {
	// Create a copy of the multiset as a slice (array), containing all values
	// len(multisetSlice) == multiset.Cardinality()
	multisetSlice := multiset.Slice()

	// Create a permutation generator - this will effectively produce all possible
	// combinations of indexes that reference values in the multiset for each variable
	// in the reaction.
	generator := combin.NewPermutationGenerator(len(multisetSlice), k)
	permutation := make([]int, k) // destination slice - the generator will populate this with it's output

	// Continuously try reactions using values determined by the permutation
	// generator until one 'happens' (the reaction condition is satisfied).
	for generator.Next() {

		// Populate the 'permutation' slice with the next permutation from the generator
		generator.Permutation(permutation)

		// Extract reactants from the multiset
		reactants := make([]ast.IntTuple, k)
		for i := 0; i < k; i++ {
			reactants[i] = multisetSlice[permutation[i]]
		}

		ok, err := performReaction(prog, k, multiset, reactants)
		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}

// Attempts to perform a reaction using the given reactants on the multiset.
// Returns true if a reaction took place, false otherwise.
func performReaction(prog *ast.Reaction, k int, multiset *Multiset, reactants []ast.IntTuple) (bool, error) {
	// Create & populate a new state to hold the program variables during the reaction
	programVariables := NewState()

	for i := 0; i < k; i++ {
		identTuple := prog.Input.Idents[i]
		valueTuple := reactants[i]

		// If the shape of the identifier tuple doesn't match the shape of the value tuple,
		// then a reaction is not possible, return false
		if !ast.ShapeMatches(identTuple, valueTuple) {
			return false, nil
		}

		for i, ident := range identTuple.Values {
			programVariables.PutVar(ident, valueTuple.Values[i])
		}
	}

	// Test the reaction condition - if it evaluates true, then a reaction can take place.
	cond, err := prog.Condition.Expression.Eval(programVariables)
	if err != nil {
		return false, errors.Wrap(err, "error evaluating reaction condition")
	}

	if !cond {
		return false, nil
	}

	// Remove the reaction inputs from the multiset
	for i := 0; i < k; i++ {
		multiset.Take(reactants[i])
	}

	// Add the reaction outputs (products) to the multiset
	for _, aexpTuple := range prog.Action.Products {
		products := make([]int, 0, aexpTuple.Dimensions())
		for _, aexp := range aexpTuple.Values {
			product, err := aexp.Eval(programVariables)
			if err != nil {
				return false, errors.Wrap(err, "error evaluating reaction product")
			}
			products = append(products, product)
		}
		multiset.Add(ast.CreateIntTuple(products))
	}

	return true, nil
}
