# Primitive Types

These are the primitive types defined. New type definition is yet to be implemented, so these are the only usable types.

## LispType
LispType is the interface that all types implement.
All types have the methods `GetType`, `ToString` and `ValueToString`.
All types also have a `GetContent` method, but since each type encapsulates a different Go type, it's not used in the interface. For example, the Boolean Type encapsulates Go's `bool` type, but the Number type encapsulates the `float64` type.


Note: `ToString` returns information of the type and the content of a variable, while `ValueToString` only returns the content.
This is used to differentiate a string for debugging and other string for output/to be used as string in other functions.
For example, we have 2 variables, a `Number` that contains the value `4.000` and a `LispString` that contains the value `"4.000"`. The method `ToString` will return "NUMBER 4.000" and "STRING 4.000", while `ValueToString` will return "4.000" for both.

### Atom (Deprecated)
Every Type in GarbageLisp is an Atom.
Might be removed in later versions.


### Boolean
Contains a Boolean type, either `true` or `false`. Go's native `nil` is also considered as `false`.

### Exp
Encapsulates Atoms/Expressions returned by the Parser.
Example: `(print (+ 1 1))` is an `Exp` that contains an Atom (`print`) and another `Exp` (`(+ 1 1)`).

### LispString
Contains a Go `string`. 
Example: `"Hey"`.

### List
Contains a List, and this List can include any other LispType, and a single List can include different ListTypes.
Example: `(list 4 5 6)` is valid, and so is `(list "hey" 4 print)`. `(quote (print 4))` will also return a list (`[print],[4.000]`).


### Number
Contains a Go `float64` value. Any text that can be parsed as a float64 by Go's `strconv.ParseFloat()` will be used as a number.
This means that `4.0` is a number, but `2A` is a symbol.
Example: `10` or `4.0`


### Procedure
Contains the Name of the Function, it's Body and it's possible Arguments/Parameters.
Also, it includes it's own Environment (or Scope), which initially includes it's parameters so they can be used to evaluate the function. 
Functions can created using the `lambda` keyword/native function, which takes a list of symbols (arguments) and an expression.
Native Functions are also stored as Procedures in the Top/Outside Scope.

Example: 
This function takes two numbers  (`(a b)`) and returns the sum of those two numbers (`(+ a b)`)
```lisp
(lambda 
    ;Argument/Parameter definition
    (a b) 
    ;Procedure Body
    (+ a b)
)
```

It's also possible to pass this Procedures as arguments.
Example:
This function takes another function/Procedure and uses the number 4 as that function argument:

```lisp
; Defining a new Procedure, "square", which takes a number as an argument and squares it
(define square 
    (lambda (a)
	(* a a)
    )
)

; Using the Map function, which takes a Procedure and List, and applies the Procedure 
; in every element of the list and returns a new list

(map square (list 1 2 3 4))

; Returns a list with [1.000],[4.000],[9.000],[16.000]

```


### Symbol
Contains a Go `string` value. It defines a symbol which is a representation of a value, this being any type that is defined in the scope.
There are some used Symbols used by Native Functions or Constants, but it's possible to make use of Symbols by defining values, using
the `define` keyword.
Example:

Defining the value of Pi.
```lisp
(define pi 3.14159265359)
```


# Native Functions Implemented


## Builtin
These functions are defined directly in the Evaluator (check Eval.go). They are not lambdas, yet they are reserved keywords.
Most of these are used to evaluate the code or are special and couldn't be implemented as "Native Functions".


### Define
Defines new value/variable

