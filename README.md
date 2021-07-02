![logo](https://i.imgur.com/KGKYp3F.png)


# Garbage Lisp  
Well it works but isn't good.
At least is **TURING COMPLETE**! 

Check the implementation of [Rule 110](https://github.com/Gabulhas/Garbage-Lisp/tree/master/examples/rule110.gl).
This automaton is Turing Complete (check [Here](https://en.wikipedia.org/wiki/Rule_110)), meaning that the implementation of it proves that the system
is also turing complete.


**Go** is definitely not the best language to do this, so it came out kinda hacky with **Go** Interfaces.

I coded this to learn about Interpreters, so this is definitely not a good example nor it works like most lisps so far.

I tried following this guide http://norvig.com/lispy.html, yet I had to change a lot of stuff because **Go** isn't Dynamically
typed like **Python** and would be easier to code this using some language that has generics. 
I'm not saying that **Go** needs them, I like it's minimalism :) .

# TODO functionality wise
- [ ] New Type Definition
- [ ] Files 
- [ ] Import other modules
- [ ] Use Linked Lists instead of Go Slices
- [ ] concurrency (jk)
- [ ] Print error location

# TODO code wise
- [ ] Document the code


# How to Run

1. Clone the Project Repo `git clone [url of this repo]`
2. Compile with Go `go build .`
3. Run with `./Garbage-Lisp`
- If you don't supply any argument, it starts a REPL.
- If you want to read from a file, you must supply the path to the file, or - if you want to read from stdin.


# Example

Fibonacci function/procedure:
```lisp
(define fib 
    ; Defining a Lambda that takes 3 arguments, "a", "b" and "endvalue"

    (lambda (a b endvalue) (
	if (> b endvalue) 
	    ; if True
	    a
	    ; else
	    (fib b (+ a b) endvalue)
	)
    )
)
```


REPL Execution: 

![execution](https://i.imgur.com/9a3uPNn.png)


If you use the `-load` flag, it will execute a file and start a REPL after it. 
This is useful to debug files by using the REPL. Example: `./Garbage-Lisp -load my_code.gl`


Check the [Examples](https://github.com/Gabulhas/Garbage-Lisp/tree/master/examples) folder.
- `42.gl`       - Big program with multiple user defined functions and input.
- `functional.gl` - Passing a function as parameter/functions as first-class citizens.
- `listCount.gl`- User defined function (lambda) that counts the number of elements equal to an atom.
- `strings.gl`  - Some basic string functions.
- `test.gl`     - Simple define and print value.
- `map.gl`      - Example implementation of the `map` keyword directly in GarbageLisp
- `metaprogramming_sortof.gl` - Sort of Metaprogramming. Using Quote and Evaluating Lists as S-expressions.
- `rule110.gl` - Implementation of Rule 110 in GarbageLisp, also proof that GarbageLisp is Turing Complete.


Check the [Examples](https://github.com/Gabulhas/Garbage-Lisp/tree/master/FUNCTIONS.md) file to get more info about the interpreter.

# Project Structure
- Env       - Environment Logic, like variables in the scope and native functions.
- Evaluator - AST walker. Includes interpretation/evaluation logic.
- LispTypes - All possible types. LispToken is the name of the Interface which all types implement.
- Parser    - Parser (obviously). Transforms code into an AST tree.
- main.go   - Starting File. Parses command line argument. Either Reads from a file (or Stdin) or starts a REPL.
