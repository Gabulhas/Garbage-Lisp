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
; to every element of the list and returns a new list

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

---

# Native Functions Implemented

---

## Builtin
These functions are defined directly in the Evaluator (check Eval.go). They are not lambdas, yet they are reserved keywords.
Most of these are used to evaluate the code or are special and couldn't be implemented as "Native Functions".


### define
Defines new value/variable in the inner most environment/scope.
Usage: `(define SYMBOL_NAME VALUE)`


Example: 
Defining a number
```lisp
(define pi 3.14159265359)
```

Defining a Procedure
```lisp
(define square 
    (lambda (a)
	(* a a)
    )
)
```


### if
Evaluates the second argument if the first argument evaluates to `true`,
evaluates the third argument otherwise.
Usage: `(if (THEN BRANCH) (ELSE BRANCH))`

Example: 
```lisp
; if 1 is less then 2
(if  (< 1 2)
     ; then
     (print "is smaller")
     ; else
     (print "is bigger")
)

;Output: "is smaller"
```


### lambda
Creates a new lambda/anonymous function.
First argument defines the names/symbols of the arguments, and the second argument defines 
de function body.
Usage: `(lambda (ARGUMENTS) (BODY))`

Example: 
Tail-Recursive Fibonacci.
```lisp
; Defining
(define fibonacci
    ; a, b and endvalue are he arguments
    (lambda (a b endvalue) 
	; Procedure body
	(if (> b endvalue) 
		; if True
		a
		; else
		(fib b (+ a b) endvalue)
	    )
    )
)
; Function execution
; calculating fibonacci(10)
(fibonacci 0 1 10)
```

### map
Calls a function on every element of an array, resulting in a new array.
Usage: `(lambda (FUNCTION/CALLBACK) (LIST))`

Example:
Defining a new function `square` and calling the function with every element of the list.
```lisp
(define square 
    (lambda (a)
	(* a a)
    )
)


(map square (list 1 2 3 4))
; (square 1)
; (square 2)
; (square 3)
; (square 4)

; Returns a list with [1.000],[4.000],[9.000],[16.000]
```


### all
Evaluates every argument (variadic arguments) and returns nil.
This function encapsulates/surrounds all the inner expressions that resulted from parsing.
Usage: `(all EXP1 EXP2 EXP3 EXP4 EXP5)`

```lisp
(all
    (print "hey")
    (print "hello")
    (print (+ 1 2))
    (define my_number 4)
    (print my_number)
)
;Output: "hey"
;Output: "hello"
;Output: "3"
;Output: "4"
; Returns nil
```

### quote
Returns the contents of a expression.
For example, this can be used to store an expression, manipulate the expression and evaluating it during execution.
Usage: `(quote EXP1)`

Example:
```lisp
(quote (print 4))

; Returns a list with [print],[4.000]
```

### Eval
```lisp
; ++ concatenates two lists
(define my_exp
    (++ 
	(quote (print 4)) 
	(list 6)
    )
)

; my_exp: a list with [print],[4.000],[6.000]


(eval my_exp)

;Output: "4"
;Output: "6"

```

As you can see in this example, we can change the expression, in this case we added a new atom to the expression `6`
and we can later evaluate it.

---
## Native
These functions are defined as primitive/native functions (NativeFunctions.go). 
Unlike builtins, these functions are lambdas defined in the outer most environment/scope.
These are implemented to enable the creation of non-native functions, making it possible to define GarbageLisp functions.


## Arithmetic functions
These functions are used to manipulate data using numbers.
If you are unfamiliar with Lisp, these arithmetic functions use a Polish notation, which means that the arithmetic
operation is defined before the numbers.
Example:
```lisp
(+ 1 2)
; returns 3, equivalent to 1 + 2

(+ 1 2 3)
; returns 6, equivalent to 1 + 2 + 3
```

INCOMPLETE
