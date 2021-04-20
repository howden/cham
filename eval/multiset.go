package eval

import (
	"fmt"
	"github.com/howden/cham/ast"
)

type Multiset struct {
	m    map[ast.IntTuple]int
	card int
}

func NewMultiset() *Multiset {
	return &Multiset{make(map[ast.IntTuple]int), 0}
}

func (set *Multiset) Add(i ast.IntTuple) {
	set.m[i]++
	set.card++
}

func (set *Multiset) AddAll(arr []ast.IntTuple) {
	for _, i := range arr {
		set.m[i]++
	}
	set.card += len(arr)
}

func (set *Multiset) Take(i ast.IntTuple) {
	existing, ok := set.m[i]
	if !ok {
		panic(fmt.Sprintf("cannot take %v from multiset", i))
	}

	if existing == 1 {
		delete(set.m, i)
	} else {
		set.m[i]--
	}
	set.card--
}

func (set *Multiset) Cardinality() int {
	return set.card
}

func (set *Multiset) Slice() []ast.IntTuple {
	res := make([]ast.IntTuple, 0, set.card)
	for val, count := range set.m {
		for i := 0; i < count; i++ {
			res = append(res, val)
		}
	}
	return res
}

func (set *Multiset) String() string {
	return fmt.Sprint(set.Slice())
}
