package eval

import (
	"github.com/howden/cham/analysis"
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
	// Complete reactions in parallel
	if analysis.DetermineReactionType(prog) == analysis.Expanding {
		err := executeParallelReactionExpanding(prog, multiset)
		if err != nil {
			return err
		}
	} else if analysis.DetermineReactionType(prog) == analysis.Shrinking {
		err := executeParallelReactionShrinking(prog, multiset)
		if err != nil {
			return err
		}
	} else { // Constant
		err := executeParallelReactionConstant(prog, multiset)
		if err != nil {
			return err
		}
	}

	// Perform a final overall reaction pass on the whole multiset.
	// This is strictly necessary for "constant" and "shrinking" reactions as the parallel execution
	// will have only partially reacted the solution.
	// For "expanding" reactions, this final pass is less likely to be necessary, but in certain edge cases,
	// there is a possibility for left-over molecules that can still react. For example if the reaction has conditions
	// depending on the relations between more than one input, some reactions may have been blocked due to partitioning.
	return performReactions(prog, multiset, -1)
}

// Parallel evaluation implementation for expanding reactions (reactions that produce more outputs than inputs).
// The general approach is to split the input multiset into partitions of size=1, then perform a single reaction on each
// partition separately in parallel, then repeat this (still in parallel) if the solution changed, and eventually
// merge the multisets back together at the end.
func executeParallelReactionExpanding(prog *ast.Reaction, multiset *Multiset) error {
	return executeParallelReaction(multiset, 1 /* partition size */, func(partition *Multiset) error {
		// First, record the starting cardinality of the partition
		before := partition.Cardinality()

		// Then, attempt to perform a single reaction 'step' on the solution
		err := performReactions(prog, partition, 1)
		if err != nil {
			return err
		}

		// If the multiset expanded as a result of the reaction, recursively call 'executeParallelReactionExpanding'
		// to partition again and repeat this process
		if partition.Cardinality() > before {
			err = executeParallelReactionExpanding(prog, partition)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// Parallel evaluation implementation for shrinking reactions (reactions that produce fewer outputs than inputs).
// The general approach is to split the input multiset into partitions of size=8, then perform reactions on each
// partition separately in parallel, increasing the partition size after each iteration by increments of 8, and
// eventually merging the multisets back together at the end.
func executeParallelReactionShrinking(prog *ast.Reaction, multiset *Multiset) error {
	partitionSize := 8

	for partitionSize*2 < multiset.Cardinality() {
		before := multiset.Cardinality()

		err := executeParallelReaction(multiset, partitionSize, func(partition *Multiset) error {
			return performReactions(prog, partition, -1)
		})
		if err != nil {
			return err
		}

		// If the multiset was unchanged by the reaction, skip further steps and return
		// This condition will be met by reactions that shrink unconditionally
		if before == multiset.Cardinality() {
			return nil
		}

		partitionSize += 8
	}

	return nil
}

// Parallel evaluation implementation for constant reactions (reactions that produce the same amount of outputs as they
// have inputs).
// The general approach is to split the input multiset into partitions of size=32, then perform reactions on each
// partition separately in parallel, then merge the multisets back together at the end.
func executeParallelReactionConstant(prog *ast.Reaction, multiset *Multiset) error {
	return executeParallelReaction(multiset, 32 /* partition size */, func(partition *Multiset) error {
		return performReactions(prog, partition, -1)
	})
}

// Generic parallel evaluation function.
// f is the function that is called to perform reactions on partitions.
func executeParallelReaction(multiset *Multiset, partitionSize int, f func(partition *Multiset) error) error {
	// Split the input multiset into partitions of the specified size
	partitions := multiset.Partition(partitionSize)

	// Create a channel to receive status callbacks from child goroutines
	c := make(chan error)

	// Iterate through each partition and schedule a goroutine to perform a parallel reaction
	for _, partition := range partitions {
		go func(partition *Multiset) {
			c <- f(partition)
		}(partition)
	}

	// Wait for all goroutines to complete
	// If any of them produce an error, return it
	n := len(partitions)
	for i := 0; i < n; i++ {
		err := <-c
		if err != nil {
			return err
		}
	}

	// Merge step: clear the original input multiset, then re-add the results of each partition
	multiset.Clear()
	for _, partition := range partitions {
		multiset.MergeFrom(partition)
	}

	return nil
}

// Performs reactions exhaustively (until no more can happen)
func performReactions(prog *ast.Reaction, multiset *Multiset, limit int) error {

	// Obtain a list of the identifiers used by the (single) reaction rule
	// The length of this array becomes 'k' in the k-permutations calculation
	k := len(prog.Input.Idents)

	// Keep track of number of reactions performed
	count := 0

	// Continuously attempt reactions until either:
	// - 'n < k' where n is the cardinality of the multiset and k is the number of identifiers
	//     i.e. there's more variables in the reaction than there are values to fill them
	// - a previous iteration of the loop was unable to complete a single reaction
	// - the number of reactions performed >= limit
	for multiset.Cardinality() >= k {
		didReactionOccur, err := attemptReaction(prog, k, multiset)
		if err != nil {
			return err
		}

		// If a reaction didn't occur (solution is "stable") then return
		if !didReactionOccur {
			return nil
		}

		// Increment reaction counter & check if limit has been reached
		count++
		if limit > 0 && count >= limit {
			return nil
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
