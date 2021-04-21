# CHAM v1.0 - Syntax Specification

This document outlines the context free grammar that describes the syntax of CHAM programs.

The code blocks contain the grammar in Backus-Naur form. The surrounding text is simply provide further explaination and is not part of the language specification.

A full (uncommented) version of the grammar is also available.

#### Contents
* [Basics](#basics)
    * [Digits and Numbers](#digits-and-numbers)
    * [Characters and Identifiers](#characters-and-identifiers)
    * [Grouping and Separators](#grouping-and-separators)
    * [Variables](#variables)
    * [Comparisons](#comparisons)
    * [Boolean Expressions](#boolean-expressions)
    * [Arithmetic Expressions](#arithmetic-expressions)
    * [Tuples](#tuples)
* [Reactions](#reactions)
    * [Reaction Input](#reaction-input)
    * [Reaction Output](#reaction-output)
    * [Reaction Condition](#reaction-condition)
    * [Reaction](#reaction)
* [Programs](#programs)
    * [Program Input](#program-input)
    * [Program](#program)


## Basics

### Digits and Numbers
```ebnf
<digit> ::= '0' | '1' | '2' | ... | '7' | '8' | '9'
<number> ::= <digit> {<digit>}
```

A `<number>` is one or more digits.

### Characters and Identifiers
```ebnf
<char> ::= 'a' | 'b' | 'c' | ... | 'x' | 'y' | 'z' | '_'
<ident> ::= <char> {<char>}
```

An `<ident>` (identifier) is one or more lower-case characters or underscores. Identifiers cannot contain upper case characters or digits.

### Grouping and Separators
```ebnf
<openb> ::= '('
<closeb> ::= ')'
<opencb> ::= '{'
<closecb> ::= '}'
<opensb> ::= '['
<closesb> ::= ']'
<comma> ::= ','
```

These tokens are used as part of other syntax constructs for either grouping, to control operator associativity or to separate distinct elements.

### Variables
```ebnf
<variable> ::= <ident> | <number>
```

A `<variable>` is something that evaluates to a number, but isn't an arithmetic expression itself. It can therefore either be an identifier (i.e. a reference to a number) or a number itself.

In a future version, we may extend the language to support more data types (floats, pairs, tuples, etc), but this version only supports integers.

> **Examples**
>
> ```
> x
> 21
> ```

### Comparisons
```ebnf
<comp-op> ::= '<' | '>' | '<=' | '>=' | '==' | '!='
<comparison> ::= <aexp> <comp-op> <aexp>
```

A `<comparison>` is an operation that takes place between two arithmetic expressions (`<aexp>`s). An arithemetic expression is just an expression that eventually evaluates to a `<variable>`.

The comparison operator (`<comp-op>`) defines the behaviour of the comparison. The comparison as a whole evaluates to a boolean.

> **Examples**
>
> ```
> a >= b
> x+1 != y
> ```

### Boolean Expressions
```ebnf
<bool-value> ::= <comparison>
```

A boolean value (`<bool-value>`) is something that evaluates to a boolean (true/false), but isn't an expression itself.

There are no constant boolean values in the language (`true` or `false`). Boolean values can only exist from the result of a comparison. Of course, `true` can be easily encoded as `0 == 0` and `false` as `0 != 0`.

```ebnf
<or> ::= '||'
<and> ::= '&&'
<not> ::= '!'

<bexp> ::= <bterm> {<or> <bterm>}
<bterm> ::= <bnotfactor> {<and> <bnotfactor>}
<bnotfactor> ::= <not> <bfactor>
<bnotfactor> ::= <bfactor>
<bfactor> ::= <bool-value>
<bfactor> ::= <openb> <bexp> <closeb>
```

A boolean expression (`<bexp>`) is an expression that combines boolean values together using logical operators (OR, AND, NOT).

> **Examples** (`<bexp>`)
>
> ```
> x <= a+b || a >= b
> !(x>1 && y>5)
> ```

### Arithmetic Expressions
```ebnf
<plus> ::= '+'
<subtract> ::= '-'
<multiply> ::= '*'
<divide> ::= '/'
<modulo> ::= '%'

<addop> ::= <plus> | <subtract>
<multop> ::= <multiply> | <divide> | <modulo>

<aexp> ::= <aterm> {<addop> <aexp>}
<aterm> ::= <afactor> {<multop> <aterm>}
<afactor> ::= <variable>
<afactor> ::= <openb> <aexp> <closeb>
```

An arithmetic expression (`<aexp>`) is an expression that combines integer values (`<variable>`s) together using arithmetic operations, producing another integer value  (a `<number>`) as output.

The supported artithmetic operators are addition, subtraction, multiplication, integer division, and modulo (remainder).

> **Examples** (`<aexp>`)
>
> ```
> x
> 5
> x+1
> 2-x
> (x+3)*z
> ```


### Tuples
Tuples are composites of other elements, denoted by square brackets (`[` `]`).
```ebnf
<ident-items> ::= <ident> {<comma> <ident>}
<ident-tuple> ::= <ident>
<ident-tuple> ::= <opensb> <ident-items> <opensb>

<number-items> ::= <number> {<comma> <number>}
<number-tuple> ::= <number>
<number-tuple> ::= <opensb> <number-items> <opensb>

<aexp-items> ::= <aexp> {<comma> <aexp>}
<aexp-tuple> ::= <aexp>
<aexp-tuple> ::= <opensb> <aexp-items> <opensb>
```

A tuple is like an array, but with a shape that is known at 'compile' time. It is basically a composite value holder for zero or more ints, int terms or identifiers.

Tuples are limited to one dimension - a tuple cannot contain another tuple (can't be nested). Tuples must always contain at least one element.

Tuple rules are defined for `ident`, `number` and `aexp`.

> **Examples** (`<ident-tuple>`)
>
> ```
> x
> [x]
> [x, y]
> ```

## Reactions
A reaction is formed of input, output and a condition.

### Reaction Input
```ebnf
<reaction-input> ::= <ident-tuple> {<comma> <ident-tuple>}
```

The input into a reaction is a comma separated list of one or more identifiers/identifier tuples.

> **Examples**
>
> ```
> x
> x, y
> [i,x], y
> ```

### Reaction Output
```ebnf
<reaction-output-items> ::= <aexp-tuple> {<comma> <aexp-tuple>}
<reaction-output> ::= <opencb> <reaction-output-items> <closecb>
<reaction-output> ::= <reaction-output-items>
<reaction-output> ::= <opencb> <closecb>
```

The output of a reaction is a comma separated list of zero or more arithmetic expressions / arithmetic expression tuples.

In the case where there are no reaction products, two curly brackets must be specified (`{}`), but otherwise, these are optional.

> **Examples**
>
> ```
> {}
> {a}
> {a+1}
> {x-1, x-2}
> {[x,0], [y,y+1]}
> ```

### Reaction Condition
```ebnf
<if> ::= 'if'
<reaction-condition> ::= <if> <bexp>
```

The condition of a reaction is the `if` keyword followed by a boolean expression.

### Reaction
```ebnf
<reaction-op> ::= '=>'
<reaction> ::= <reaction-input> <reaction-op> <reaction-output>
<reaction> ::= <reaction-input> <reaction-op> <reaction-output> <reaction-condition>
```

We combine the other elements together to create a single reaction. Note that the reaction condition is optional, but all other parts are required.

> **Example**
>
> ```
>  x, y => x if x > y
> |----|               <-- reaction-input
>      |--|            <-- reaction-op
>         |-|          <-- reaction-output
>           |--------| <-- reaction-condition
> ```


## Programs
Programs are formed of an initial input multiset, followed by one or more reactions.

### Program Input
```ebnf
<program-input-items> ::= <number-tuple> {<comma> <number-tuple>}
<program-input> ::= <opencb> <program-input-items> <closecb>
<program-input> ::= <program-input-items>
```

The input into a program is a comma separated list of one or more numbers/number tuples (a multiset).

> **Examples**
>
> ```
> 1, 2, 3
> {1, 2, 3}
> {[0, 1], [0, 2], [0, 3]}
> ```

### Reaction definitions
```ebnf
<reaction-chain> ::= '|'

<reaction-def-operator> ::= ':'
<reaction-def-statement> ::= <ident> <reaction-def-operator> <reaction> {<reaction-chain> <reaction>}
```

When in REPL mode, it is possible to define and store reactions for later use.

The `<reaction-chain>` operator is used to link reactions together.

> **Example**: single reaction (`<reaction-def-statement>`)
>
> ```
>  max : x, y => x if x > y
> |---|                      <-- ident
>     |-|                    <-- reaction-def-operator
>       |------------------| <-- reaction
> ```
>
> **Example**: multiple reactions (`<reaction-def-statement>`)
>
> ```
>  fib : x => {x-1, x-2} if x>1 | x,y => x+y
> |---|                                       <-- ident
>     |-|                                     <-- reaction-def-operator
>       |----------------------|              <-- reaction
>                              |-|            <-- reaction-chain
>                                |----------| <-- reaction
> ```


### Program
```ebnf
<reaction-pointer> ::= <reaction>
<reaction-pointer> ::= <reaction-def-operator> <ident>

<program> ::= <program-input> <reaction-chain> <reaction-pointer> {<reaction-chain> <reaction-pointer>}
```

A program is made up of initial input followed by an executable chain of reactions (or reaction pointers).

Reaction pointers are available in REPL mode, and point to reactions which have already been defined.

> **Example**: single reaction
>
> ```
>  {1,2,4,7,3} | x, y => x if x > y
> |-----------|                      <-- program-input
>             |-|                    <-- reaction-chain
>               |------------------| <-- reaction
> ```
>
> **Example**: multiple reactions
>
> ```
>  {7} | x => {x-1, x-2} if x>1 | x,y => x+y
> |---|                                       <-- program-input
>     |-|                                     <-- reaction-chain
>       |----------------------|              <-- reaction
>                              |-|            <-- reaction-chain
>                                |----------| <-- reaction
> ```
