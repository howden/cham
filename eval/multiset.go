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

func (set *Multiset) MergeFrom(other *Multiset) {
	for k, v := range other.m {
		set.m[k] += v
	}
	set.card += other.card
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

func (set *Multiset) Clear() {
	set.m = make(map[ast.IntTuple]int)
	set.card = 0
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

// Partitions the multiset into multiple other multisets of the given size
func (set *Multiset) Partition(size int) []*Multiset {
	if size <= 0 {
		panic("size cannot be <= 0")
	}

	slice := set.Slice()
	length := len(slice)

	if length <= 0 {
		return []*Multiset{}
	}

	n := length / size
	res := make([]*Multiset, 0, n+1)

	f := func(lo int, hi int) {
		newSet := NewMultiset()
		for _, e := range slice[lo:hi] {
			newSet.Add(e)
		}
		res = append(res, newSet)
	}

	var i int
	for i = 0; i < n; i++ {
		f(i*size, i*size+size)
	}
	if rest := length % size; rest != 0 {
		f(i*size, i*size+rest)
	}

	return res
}

func (set *Multiset) String() string {
	return fmt.Sprint(set.Slice())
}
