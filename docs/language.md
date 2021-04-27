# CHAM v1.0 - Language Guide

This document contains a brief introductory guide explaining how the language works.

#### Contents

* [The background](#the-background)
* [All programs start with data](#all-programs-start-with-data)
* [The core of everything: reactions](#the-core-of-everything-reactions)
* [Performing arithmetic](#performing-arithmetic)
* [Boolean logic in conditions](#boolean-logic-in-conditions)
* [Producing more than one output](#producing-more-than-one-output)
* [Some parts of a program are optional](#some-parts-of-a-program-are-optional)
* [Using the REPL's memory](#using-the-repls-memory)
* [Chaining reactions together](#chaining-reactions-together)
* [Tuples!](#tuples)
* [What's next?](#whats-next)

___

### The background

The language is based on the "chemical programming" paradigm used in the *Chemical Abstract Machine* (proposed by Gérard Berry and Gérard Boudol in 1992), which itself is based on the *Gamma model of programming* (proposed by Jean-Pierre Banâtre and Daniel Le Métayer in 1986). It is in no way a new idea!

In this model, programs are formed of **reaction rules**, which can be (optionally) chained together along with some initial input to make a complete program that can be executed.


### All programs start with data

All programs start with some initial input data.

For example:

```
{1,2,4,7,3}
```

The input is a *multiset*, known in the chemical paradigm as the **solution** - think of it like a "bag of numbers", with duplicates *allowed*!


### The core of everything: reactions

So we have some input data sorted.

Next, as an example, let's imagine we want to write a program to find the largest value.

From a glance, we can see that the largest number is `7`, but we need to write a program - or more specifically, a reaction - that can compute this for us for any solution.

#### General thinking

Going back to the "bag of numbers" idea, if you started with a bag full of 1000 pieces of paper, each with a number written on it, it would be *impossible* to immediately find the largest at a glance. You need to be systematic! What would you do?

One approach is:

* Take two pieces of paper out of the bag, it doesn't matter which ones - take any!
* Compare the two pieces of paper, discard the smallest one, and put the largest back into the bag
* Repeat!

At the end, you would be left with only one piece of paper in the bag - the largest number. This is (pretty much!) exactly how reactions work in the language.

#### Turning it into code

This process can be "codifed" as the following reaction:

```
x, y => { x } if x > y
```

What this code means is:

* Pick two elements, `x` and `y` out of the solution
* Providing they match the condition `x > y` (is x larger than y?)...
* Then take *both* of them out of the solution, and put `x` back in

Generally, we can see that reactions are formed of:

* A list of input variables, separated by `,` commas
* The "reaction operator" `=>`
* The reaction outputs, enclosed by curly brackets `{ }`
* The reaction condition, the keyword `if` followed by a condition

#### Running it

If we put the initial input together with the reaction, we have a program that can be executed in the interpreter!

```
{1,2,4,7,3} | x, y => { x } if x > y
```

The only thing that we've added is the `|` character between the input and the first reaction. This is called the "reaction chaining" operator. You can use it to chain reactions together.


### Performing arithmetic

You don't *just* have to return variables as-is. You can also perform arithmetic on them!

For example:

```
x => { x+5 } if x < 5
```


### Boolean logic in conditions

Conditions can contain boolean logic.

* Boolean AND: `&&`
* Boolean OR: `||`
* Boolean NOT: `!`

For example:

```
x, y, z => { x+y-z } if x>2 || (y>3 && z==y)
```


### Producing more than one output

Reactions can result in more than one item of output. For example:

```
x => { x-1, x-2 } if x > 1
```


### Some parts of a program are optional

Reactions don't necessarily need to contain any outputs! In this case, just specify the empty set `{ }` as output. For example:

```
x => { } if x > 5
```

Reactions also don't necessarily need to contain a condition! For example:

``` 
x, y => { x+y }
```


### Using the REPL's memory

Computers are (sometimes) great at remembering things. We can use this to avoid having to type so much!

You can store a program in the REPL memory using an identifier of your choosing, then reference it later.

Taking our "find the largest value" program from earlier, we can give it a name (`max`), and ask the REPL to store it for use later.

```
max: x, y => { x } if x > y
```

We can then easily "import" it in subsequent programs:

```
{1,2,4,7,3} | :max
```

Just remember to prefix the name with the `:` character!


### Chaining reactions together

You can use the `|` reaction chain operator to join reactions together.

The way this works is that the *output* of a prior reaction becomes the *input* for the next! Simple as that.


### Tuples!

As well as plain integers, solutions can contain tuples. You can think of tuples like one more more numbers organised together in a bubble.

Tuples are defined using square brackets `[]`.

For example:

```
largest_first: [x,y] => { [y,x] } if y > x 
```

Tuples allow more interesting programs to be implemented - you can check out the example programs to see!


### What's next?

The simple answer is: start playing!

If you're interested in finding out more about the language syntax, more information can be found in [syntax.md](syntax.md).

If you want to take a look at some example programs, those can be found in [programs.md](programs.md).
