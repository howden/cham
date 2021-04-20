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
	idents := prog.Input.Idents
	k := len(idents)

	// Continuously attempt reactions until either:
	// - 'n < k' where n is the cardinality of the multiset and k is the number of identifiers
	//     i.e. there's more variables in the reaction than there are values to fill them
	// - a previous iteration of the loop was unable to complete a single reaction
	solved := false
	for !solved && multiset.Cardinality() >= k {

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
		reaction := false

	permutationLoop:
		for !reaction && generator.Next() {

			// Populate the 'permutation' slice with the next permutation from the generator
			generator.Permutation(permutation)

			// Create & populate a new state to hold the program variables during the reaction
			programVariables := NewState()
			for i, identTuple := range idents {
				valueTuple := multisetSlice[permutation[i]]

				// If the shape of the identifer tuple doesn't match the shape of the value tuple,
				// then a reaction is not possible, so continue the outer loop.
				if !ast.ShapeMatches(identTuple, valueTuple) {
					continue permutationLoop
				}

				for i, ident := range identTuple.Values {
					programVariables.PutVar(ident, valueTuple.Values[i])
				}
			}

			// Test the reaction condition - if it evaluates true, then a reaction can take place.
			cond, err := prog.Condition.Expression.Eval(programVariables)
			if err != nil {
				return errors.Wrap(err, "error evaluating reaction condition")
			}

			if cond {
				reaction = true

				// Remove the reaction inputs from the multiset
				for i := 0; i < k; i++ {
					oldValue := multisetSlice[permutation[i]]
					multiset.Take(oldValue)
				}

				// Add the reaction outputs (products) to the multiset
				for _, aexpTuple := range prog.Action.Products {
					products := make([]int, 0, aexpTuple.Dimensions())
					for _, aexp := range aexpTuple.Values {
						product, err := aexp.Eval(programVariables)
						if err != nil {
							return errors.Wrap(err, "error evaluating reaction product")
						}
						products = append(products, product)
					}
					multiset.Add(ast.CreateIntTuple(products))
				}
			}
		}

		// If a reaction didn't take place during this iteration, then set solved=true
		if !reaction {
			solved = true
		}
	}

	return nil
}
