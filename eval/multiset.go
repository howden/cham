package eval

import "fmt"

type Multiset struct {
	m    map[int]int
	card int
}

func NewMultiset() *Multiset {
	return &Multiset{make(map[int]int), 0}
}

func (set *Multiset) Add(i int) {
	set.m[i]++
	set.card++
}

func (set *Multiset) AddAll(arr []int) {
	for _, i := range arr {
		set.m[i]++
	}
	set.card += len(arr)
}

func (set *Multiset) Take(i int) {
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

func (set *Multiset) Slice() []int {
	res := make([]int, 0, set.card)
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
