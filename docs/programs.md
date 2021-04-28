# Example Programs

This document contains a number of programs written using the language.

They can be tested / executed using the REPL program.

Plain, undocumented versions of these programs can be found in [programs.txt](programs.txt). You can load all of these directly into the REPL using the command `:load docs/programs.txt`.

[References](#references) are provided for programs which were derived from other sources.

#### Contents
* [`max`](#max)
* [`min`](#min)
* [`remove_duplicates`](#remove_duplicates)
* [`sum`](#sum)
* [`product`](#product)
* [`filter_odd`](#filter_odd)
* [`filter_even`](#filter_even)
* [`prime_sieve`](#prime_sieve)
* [`fib`](#fib)
* [`sort`](#sort)
* [`iota`](#iota)
* [`primes`](#primes)
* [`factorial`](#factorial)
* [`max_segment_sum`](#max_segment_sum)
* [`prime_factorization`](#prime_factorization)

___

### `max`
Filters the multiset so that only the largest remains.
```
max: x,y => x if x>y
```

> **Example**
>
> ```
> > {1,2,4,7,3} | :max
> [7]
> ```

### `min`
Filters the multiset so that only the smallest remains.
```
min: x,y => x if x<y
```

> **Example**
>
> ```
> > {1,2,4,7,3} | :max
> [1]
> ```

### `remove_duplicates`
Filters the multiset so that duplicates are removed.
```
remove_duplicates: x,y => x if x==y
```

> **Example**
>
> ```
> > {1,7,2,1,7,1} | :rem_duplicates
> [1 7 2]
> ```

### `sum`
Reduces the multiset to the sum of all elements.
```
sum: x,y => x+y
```

> **Example**
>
> ```
> > {1,7,2,4,7,3} | :sum
> [24]
> ```

### `product`
Reduces the multiset to the product of all elements.
```
product: x,y => x*y
```

> **Example**
>
> ```
> > {1,7,2,4,7,3} | :product
> [1176]
> ```

### `filter_odd`
Filters the multiset so that only odd elements remain.
```
filter_odd: x => {} if x%2 == 0
```

> **Example**
>
> ```
> > {1,7,2,4,7,3} | :filter_odd
> [3 1 7 7]
> ```

### `filter_even`
Filters the multiset so that only even elements remain.
```
filter_even: x => {} if x%2 == 1
```

> **Example**
>
> ```
> > {1,7,2,4,7,3} | :filter_even
> [2 4]
> ```

### `prime_sieve`
Given a multiset {`2..n`}, filters so only the prime numbers remain.  
Also known as the Sieve of Eratosthenes.

See https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes for more information.
```
prime_sieve: x,y => y if x%y == 0
```
[[1]](#references)
> **Example**
>
> ```
> > {2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20} | :prime_sieve
> [2 17 3 13 5 19 7 11]
> ```

### `fib`
Given a multiset {`n`}, returns the nth fibonacci number.

See https://en.wikipedia.org/wiki/Fibonacci_number for more information.
```
fib: x => { x-1, x-2 } if x>1  |  x,y => x+y
```
[[2]](#references)
> **Example**
>
> ```
> > {14} | :fib
> [377]
> ```

### `sort`
Given a multiset M of elements, returns a multiset of tuples `[i, m]` where `i` is the sorted position of element `m` in the imaginary ordering of M.
```
sort_init: x => [0,x]
sort_promote: [i,x], [j,y] => { [i+1,x], [j,y] } if i==j && x>=y
sort: :sort_init | :sort_promote
```

It is also possible to sort an existing multiset of tuples.
```
sort_existing: [i,x], [j,y] => { [i,y], [j,x] } if i<j and x>y
```
[[3]](#references)
> **Example**
>
> ```
> > {7,8,9,1,2,3} | :sort
> [[1 2] [0 1] [5 9] [4 8] [3 7] [2 3]]
> ```
> manually displayed in ascending order - remember the tuple is of the form `[index, value]`
> 
> > [[0 **1**] [1 **2**] [2 **3**] [3 **7**] [4 **8**] [5 **9**]]

### `iota`
Given a multiset {`[x,y]`}, returns a multiset of consecutive numbers in the range `x..y`.
```
iota: [x,y] => { [x, (x+y)/2], [(x+y)/2+1, y] } if x != y  |  [x,y] => x if x==y
```
[[2]](#references)
> **Example**
>
> ```
> > {[2,8]} | :iota
> [2 6 4 3 8 7 5]
> ```
> 
> Remember: multisets do not have any defined order, so the output may vary.

### `primes`
Given a multiset {`n`}, returns a multiset of all prime numbers up to `n`.  
Also known as the Sieve of Eratosthenes.

See https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes for more information.
```
primes: x => [2,x]  |  :iota  |  :prime_sieve
```
[[2]](#references)
> **Example**
>
> ```
> > {30} | :primes
> [5 23 13 2 17 7 19 29 3 11]
> ```

### `factorial`
Given a multiset {`n`}, returns a multiset containing `n!`.

See https://en.wikipedia.org/wiki/Factorial for more information.
```
factorial: x => [1,x]  |  :iota  |  :product
```
[[3]](#references)
> **Example**
>
> ```
> > {9} | :factorial
> [362880]
> ```

### `max_segment_sum`
Given an input sequence `a..b` of numbers in the format `[i,a]..[j,b]` where `i..j` corresponds to the position/index of elements `a..b` in the sequence, returns a tuple `[s,i]` where `s` is the sum and `i` is the end index of the largest segment sum.

See https://en.wikipedia.org/wiki/Maximum_subarray_problem for more information.
```
max_segment_sum: [i,x] => [i,x,x]  |
  [i,x,s], [ip,xp,sp] => [i,x,s], [ip,xp,s+xp] if ip == i+1 && s+xp > sp  |
  [i,x,s], [ip,xp,sp] => [ip,xp,sp] if sp > s  |
  [i,x,s] => [s,i]
```
[[2]](#references)
> **Example**
> 
> Using sequence `-10, -2, 1, -3, 4, -1, 2, 1, -5, 4` as an example:
> 
> ```
> > {[0,-10], [1,-2], [2,1], [3,-3], [4,4], [5,-1], [6,2], [7,1], [8,-5], [9,4]} | :max_segment_sum
> [[6 7]]
> ```
> 
> The contiguous subsequence with the largest sum is `4, -1, 2, 1` which has sum `6`, and ends at index `7`.

### `prime_factorization`
Given an input `n`, decompose n into its prime factors.

See https://en.wikipedia.org/wiki/Prime_number#Unique_factorization for more information.

```
prime_factorization_coeff: [n, p, k] => [n, p/n, k+1] if p % n == 0  |
  [n, p, k] => [n, k]  |
  [n, k] => {} if k == 0

prime_factorization_reduce: [n, k] => { [n, k-1], [n, 1] } if k > 1  |  [n, k] => n

prime_factorization:  :prime_factorization_coeff | :prime_factorization_reduce
```
[[3]](#references)
> **Example**
>
> Using the number `1925` as an example:
>
> ```
> > {1925} | :primes | n => [n, 1925, 0] | :prime_factorization
> [11 7 5 5]
> ```
>
> `1925 = 7 * 11 * 5^2`

___

### References

> [1] J.-P. Banâtre, P. Fradet, and Y. Radenac, ‘Principles of Chemical Programming’, Electronic Notes in Theoretical Computer Science, vol. 124, no. 1, Art. no. 1, Mar. 2005, doi: 10.1016/j.entcs.2004.07.019.

> [2] J.-P. Banâtre, P. Fradet, and D. Le Métayer, ‘Gamma and the chemical reaction model: fifteen years after’, p. 27, Aug. 2000, doi: 10.1142/9781848161023_0001.

> [3] J.-P. Banâtre and D. Le Métayer, ‘Programming by multiset transformation’, Commun. ACM, vol. 36, no. 1, Art. no. 1, Apr. 1990, doi: 10.1145/151233.151242.
