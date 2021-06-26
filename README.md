![logo](https://i.imgur.com/KGKYp3F.png)


# Garbage Lisp  
Well it works but isn't good.

**Go** is definitely not the best language to do this, so it came out kinda hacky with **Go** Interfaces.

I coded this to learn about Interpreters, so this is definitely not a good example nor it works like most lisps so far.

I tried following this guide http://norvig.com/lispy.html, yet I had to change a lot of stuff because **Go** isn't Dynamically
typed like **Python** and would be easier to code this using some language that has generics. 
I'm not saying that **Go** needs them, I like it's minimalism :) .

# TODO functionality wise
- [ ] Better Scoping
- [ ] New Type Definition

- [ ] "Debug Mode"
- [ ] Files 
- [ ] Import other files
- [ ] Improve Prints 
- [ ] Use Linked Lists instead of Go Slices
- [ ] concurrency (jk)
- [ ] Improve REPL
- [ ] Print error location

# TODO code wise
- [ ] Document the code
- [ ] Document native functions



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


Check the [Examples](https://github.com/Gabulhas/Garbage-Lisp/tree/master/examples) folder.
- `42.gl`       - Big program with multiple user defined functions and input.
- `functional.gl` - Passing a function as parameter/functions as first-class citizens.
- `listCount.gl`- User defined function (lambda) that counts the number of elements equal to an atom.
- `metaprogramming_sortof.gl` - Sort of Metaprogramming. Using Quote and Evaluating Lists as S-expressions.
- `strings.gl`  - Some basic string functions.
- `test.gl`     - Simple define and print value.

# Project Structure
- Env       - Environment Logic, like variables in the scope and native functions.
- Evaluator - AST walker. Includes interpretation/evaluation logic.
- LispTypes - All possible types. LispToken is the name of the Interface which all types implement.
- Parser    - Parser (obviously). Transforms code into an AST tree.
- main.go   - Starting File. Parses command line argument. Either Reads from a file (or Stdin) or starts a REPL.
